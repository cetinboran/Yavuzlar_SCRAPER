# Yavuzlar_SCRAPER
+ This project serves the purpose of extracting desired information from websites. It is a scraper project written in Go.

# How to Install?
+ You can download it by running `go get github.com/cetinboran/scrapergo`.

# Basic Structure
+ Send a request to the desired website and obtain the response body.
+ Pass the obtained body to the scraper.BodyReader(responseBody) function. This function returns two arguments: a scraper object and an error object.
+ You can perform searches using this scraper object. Now, let's discuss the functions.

# Structs
+ Config:
    + `AutoSave bool`: If false, it will not save automatically. If true, it saves automatically
+ Tag:
    + `SetName()`: You add the tag name with this.
    + `SetClasses()`: You add class to the tag. Comma separator.
    + `SetId()`: You can add the id to the tag.
    + `SetAttiributes()`: You can add the id to the tag. Comma separator.
+ Collector:
    + `Each(func (i int, name string){})`: Loops through the data you found.
    + `GetData()`: Returns the data you found.
    + `Save()`: Saves the data into collection.json. You cannot use this function if AutoSave is true.

# TagStr Format
+ When a function requests a tagStr from you, you should enter it in the following format:
    + If you write it plainly, it will be added as a tag name.
        + div
        + form
    + If you start with . it adds it as a class
        + .title
        + .title .article
    + if you start with # it takes as id
    + More than one class can be added like "title article" but only one id can be added.
    + If you write between [] you add a attribute.
        + [href]
        + [value]

# Scraper Functions
+ `scraper.Use(c Config)`:
    + You can add any config settings you want with this function. If you don't use the default config will be used
+ `scraper.FindWithTag(t Tag)`:
    + This function takes a tag object and performs a search based on the tag name, classes, and ID added to this tag object.
+ `scraper.Find(tagStr string)`:
    + This function takes a tagStr string and search based on that.
+ `scraper.FindAttr(tagStr string, attr string)`:
    + This function takes the first parameter as tagStr, and in the second parameter, you specify the attribute you want to find. For example, if you write "a [href]", it will search for tags with the attribute "href" within the "a" tags. If you specify "href" as the second parameter, it will return the values of the "href" attribute found in the tags it discovers.
+ `scraper.FindWithRegex(tagStr string, regex string)`:
    + This function takes the first parameter as tagStr, and the second parameter is a regular expression (regex). Once a tag is created based on the first argument, it searches within these tags for values that match the provided regex in the second argument. For example, if you input 'body' and '\d{11}' as arguments, it will return any sequences of 11 consecutive digits found within the body tag.
+ `scraper.FindLinks(), scraper.FindEmails()`:
    + This is a shortcut method that provides ready-made functionality for finding links within the content. You can achieve similar results using the functions mentioned above, but this one is specifically designed for conveniently finding links.



# Contact
<p align="center">
  <a href="https://github.com/cetinboran">
    <img src="https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg" alt="github" height="40">
  </a>
  <a href="https://www.linkedin.com/in/cetinboran-mesum/">
    <img src="https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg" alt="linkedin" height="40">
  </a>
  <a href="https://www.instagram.com/2023an_m/">
    <img src="https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/instagram.svg" alt="instagram" height="40">
  </a>
  <a href="https://twitter.com/2023anM">
    <img src="https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/twitter.svg" alt="twitter" height="40">
  </a>
</p>