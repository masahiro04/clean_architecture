package testData

import (
	"clean_architecture/golang/domains"
	blog "clean_architecture/golang/domains/blog"
	"strconv"
	"time"
)

func Blog() domains.Blog {
	id, _ := blog.NewId(1)
	title, _ := blog.NewTitle("タイトル")
	body, _ := blog.NewBody("内容")
	return domains.BuildBlog(id, title, body, time.Time{}, time.Time{})
}

func Blogs(length int) domains.Blogs {
	var blogs []domains.Blog
	n := 1
	for n < length {
		id, _ := blog.NewId(n)
		title, _ := blog.NewTitle("タイトル" + strconv.Itoa(n))
		body, _ := blog.NewBody("内容" + strconv.Itoa(n))
		blogs = append(blogs, domains.BuildBlog(id, title, body, time.Time{}, time.Time{}))
		n++
	}
	return domains.NewBlogs(blogs)
}
