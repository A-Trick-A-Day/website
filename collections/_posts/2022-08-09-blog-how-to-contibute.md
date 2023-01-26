---
title: "How to contribute"
date: 2022-08-09T07:49:03+01:00
layout: post
authors: ["Sidorenko Konstantin"]
categories: ["Blog"]
description: Want to write posts, but don't know how, check out this post.
thumbnail: "assets/images/thumbnail/2022-08-09-blog-how-to-contibute.jpg"
comments: true
---

## Contact US

**First**, contact the [main contributors]({{ site.data.social[0].link }}/graphs/contributors) to discuss the **topic you want to talk** about.

After validation, make a [fork of the github project]({{ site.data.social[0].link }}/fork) to create your future merger request.

{% include alerts/info.html content='For future contributions, you will surely be added to the project so no need to fork.' %}

## Create you first Post

Checkout the [README.md install section]({{ site.data.social[0].link }}#installing-theme), to install everything.

When all is okey, **edit** the following files:

```bash
.
├── assets
│   ├── images
│   │   ├── author
│   │   │   └── sidorenko-konstantin.jpg # 2
│   │   └── thumbnail # 5
│   │       ├── 2022-08-03-golang-anonymous-functions.jpg
│   │       ├── 2022-08-04-golang-project-structure.jpg
│   │       └── 2022-08-08-golang-overview-of-the-go-tooling.png
├── categories # 4
│   ├── ci.md
│   ├── development.md
│   ├── golang.md
│   └── tool.md
├── collections
│   └── _posts # 3
│       ├── 2022-08-03-golang-anonymous-functions.md
│       ├── 2022-08-04-golang-project-structure.md
│       └── 2022-08-08-golang-overview-of-the-go-tooling.md
└── _data
    ├── authors.yml # 1
    ├── menu.yml
    └── social.json
```

1. First of all, if you are a **new author**, add yourself to the list.

   - Example of author:

   ```yml
   example_name:
     name: "Example Name"
     image: assets/images/author/example-name.{png,jpg}
     github: example
     twitter: example # not mandatory
   ```

1. Then add your most **beautiful picture**
1. You are finally ready to **write your post**, start your **file name** with the **publicationDate-thenTheMainCategory-andFinallyThePostTitle.md**.

   - Add this on the top of the markdown file

   ```md
   ---
   title: "Your main title"
   date: 2022-08-08T07:49:03+01:00 # on local development change this date to see the post but don't forget to change it
   layout: post
   authors: ["Your author name 1", "Your author name 2"]
   categories: ["Your category 1", "Your category 2"]
   description: Your description
   thumbnail: "assets/images/thumbnail/your-post-file-name.{jpg,png}"
   comments: true # if you want to enable disqus to have comment section on the post
   ---
   ```

1. If your post contains a new category add a file in the **categories** folder

   - Add this on the top of the markdown file

   ```md
   ---
   layout: category
   title: My new category
   ---
   ```

1. Add a 16:9 thumbnail to your article, name it like your post file.
   {% include alerts/info.html content='I use <a href="https://pixlr.com/fr/x/#editor">pixlr</a> to edit the thumbnails directly from a browser.' %}
   {% include alerts/info.html content='I use <a href="https://carbon.now.sh">carbon.now.sh</a> to generate code images.' %}
   {% include alerts/green.html content='Try to optimize its size to 640/320px. And your can use <strong>optipng</strong> and <strong>jpegoptim</strong> to optimize the image.' %}

Finally, **create your merger request**, for review.

**Welcome** as a new author and **thank you** for your contributions!
