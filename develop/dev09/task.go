package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Получаем массив всех ссылок сайта
func LinkParser(url string) []string {
	//Загружаем ссылку
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//обрабатываем чтобы можно было делать поиск по тегам
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//срез для результатов поиска ссылок
	var links []string

	//Ищем и записываем все ссылки на странице
	doc.Find("body a, link,head").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		links = append(links, link)
	})
	return links
}

//функция скачивания документа
func downloadDocument(path string, url string) error {
	//сплитим для выделения названия и пути сохранения
	filepath := strings.Split(url, "/")
	if len(filepath[len(filepath)-1]) == 0 {
		return nil
	}
	//пытаемся скачать абсолютную ссылку
	response, err := http.Get(url)
	if err != nil {
		//относительную
		response, err = http.Get("https://" + path + "/" + url)
		if err != nil {
			return err
		}

	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil
	}
	temp := "download/" + path + "/"
	//склеиваем пусть сохранения
	for i := 3; i < len(filepath)-2; i++ {
		temp += filepath[i] + "/"
	}
	//создаем директорию
	err = os.MkdirAll(temp, 0777)
	if err != nil {
		fmt.Println(err)
	}

	//создаем файл
	file, err := os.Create(temp + filepath[len(filepath)-1])
	if err != nil {
		return err
	}

	defer file.Close()

	//и копируем туда
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

//Скачиваем все файлы с сайта
func downloadFromPage(path string, url string) {
	err := downloadDocument(path, url)
	if err != nil {
		log.Fatal(err)
	}
	//Получаем срез ссылок
	links := LinkParser(url)
	//Для каждой из них создаем свой файл
	for _, l := range links {
		err = downloadDocument(path, l)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func main() {
	//Флаги для урла и глубины загрузки
	url := flag.String("url", "", "url")

	flag.Parse()
	//Создаем путь для загрузки страницы
	pathArr := strings.Split(*url, "/")

	downloadFromPage(pathArr[2], *url)
}
