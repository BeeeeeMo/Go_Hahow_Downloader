package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"hahow_downloader/hahow"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/TruthHun/html2md"
	log "github.com/sirupsen/logrus"
)

var CONCURRENCY_LIMIT int
var TOKEN string
var COURSE_ID string

func fetch(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return ""
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	req.Header.Set("authorization", TOKEN)
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return ""
	}

	if resp.StatusCode != 200 {
		fmt.Println("Status code error: ", resp.StatusCode)
		return ""
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return ""
	}
	bodyString := string(bodyBytes)
	return bodyString
}

func fetchTitle(courseId string) string {
	log.Debug("Fetching title: ", courseId)
	web := fetch("https://api.hahow.in/api/courses/" + courseId)
	var c = hahow.CoueseInfo{}

	// parsing json
	err := json.Unmarshal([]byte(web), &c)
	if err != nil {
		log.Error(err)
		return ""
	}
	return c.Title
}

func fetchCourse(courseId string) hahow.Course {
	log.Debug("Fetching course: ", courseId)
	web := fetch("https://api.hahow.in/api/courses/" + courseId + "/modules/items")
	var course = hahow.Course{}

	// parsing json
	err := json.Unmarshal([]byte(web), &course)
	if err != nil {
		log.Error(err)
		return hahow.Course{}
	}
	return course
}

func fetchLecture(lectureId string) hahow.Lecture {
	log.Debug("Fetching lecture: ", lectureId)
	web := fetch("https://api.hahow.in/api/lectures/" + lectureId)
	var lecture = hahow.Lecture{}

	// parsing json
	err := json.Unmarshal([]byte(web), &lecture)
	if err != nil {
		log.Error(err)
		return hahow.Lecture{}
	}
	return lecture
}

func getHighestQualityVideo(lecture hahow.Lecture) string {
	var url string
	var highestSize int
	for _, video := range lecture.Video.Videos {
		if video.Size > highestSize {
			highestSize = video.Size
			url = video.Link
		}
	}
	return url
}

func downloadLecture(lectureId, dir, chapterNumber, fileName string) {
	lecture := fetchLecture(lectureId)
	url := getHighestQualityVideo(lecture)
	log.Info("Downloaded: ", fileName)
	srtFileName := dir + "/" + fileName + ".srt"
	if _, err := os.Stat(srtFileName); os.IsNotExist(err) {
		downloadFile(url, dir+"/"+fileName+".mp4")
		for _, v := range lecture.Video.Subtitles {
			if v.Language == "zh-TW" {
				downloadFile(v.Link, srtFileName)
			}
		}
	}

	mdFileName := dir + "/" + chapterNumber + "_老師的話.md"
	if lecture.Description != "" {
		log.Debug("Writing description file: ", mdFileName)
		md := html2md.Convert(lecture.Description)
		ioutil.WriteFile(mdFileName, []byte(md), 0644)
	}

	if len(lecture.Assets) > 0 {
		var data string
		assetsFileName := dir + "/" + chapterNumber + "_附件.txt"
		log.Debug("Writing assets file: ", assetsFileName)
		if _, err := os.Stat(assetsFileName); os.IsNotExist(err) {
			for _, v := range lecture.Assets {
				data += v.DisplayName + "\r\n" + v.URL + "\r\n\r\n"
			}
			ioutil.WriteFile(assetsFileName, []byte(data), 0644)
		}
	}

}

func downloadFile(url string, path string) {
	log.Debug("Downloading file: ", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	req.Header.Set("authorization", TOKEN)
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		log.Error(err)
		return
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
}

func cleanName(Name string) string {
	regex := regexp.MustCompile(`[\\/:*?"<>|]`)
	return regex.ReplaceAllString(Name, "")
}

func mkdir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// if config file exists, read token from config file
	if _, err := os.Stat("config.json"); err == nil {
		file, _ := ioutil.ReadFile("config.json")
		var config = hahow.Config{}
		err := json.Unmarshal([]byte(file), &config)
		if err != nil {
			log.Error(err)
			return
		}
		TOKEN = config.Token
		COURSE_ID = config.CourseID
		CONCURRENCY_LIMIT = config.ConcurrencyLimit
	} else {
		// read token from args -T
		if len(os.Args) < 2 {
			// print help info
			fmt.Println("--help for more info")
			os.Exit(1)
		}

		flag.StringVar(&TOKEN, "T", "", "Token")
		// read course id from args -C
		flag.StringVar(&COURSE_ID, "C", "", "Course ID")
		// read concurrency limit from args -L defalut 2
		flag.IntVar(&CONCURRENCY_LIMIT, "L", 2, "Concurrency Limit")
		flag.Parse()
	}
}

func main() {
	// start time
	start := time.Now()
	courseTitle := cleanName(fetchTitle(COURSE_ID))
	log.Info("Course Downloading: ", courseTitle)
	var wg sync.WaitGroup
	ch := make(chan struct{}, CONCURRENCY_LIMIT)
	Course := fetchCourse(COURSE_ID)
	for _, chapter := range Course {
		chapterTitle := cleanName(chapter.ChapterNumber + "_" + chapter.Title)
		for _, lecture := range chapter.Items {
			lectureTitle := cleanName(lecture.Content.Title)
			if lecture.Type != "LECTURE" {
				continue
			}
			filePath := "./" + courseTitle + "/" + chapterTitle
			fileName := lecture.ChapterNumber + "_" + lectureTitle
			wg.Add(1)
			// Download lecture Video
			go func(lectureId, dir, fileName, chapterNumber string) {
				ch <- struct{}{}
				defer func() {
					<-ch
					wg.Done()
				}()
				mkdir(filePath)
				downloadLecture(lectureId, dir, chapterNumber, fileName)
			}(lecture.Content.ID, filePath, fileName, lecture.ChapterNumber)

		}
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Info("Used time: ", elapsed)
}
