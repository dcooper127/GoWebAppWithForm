<html>
    <head>
        <style>
            div.myText{max-width:40%;}
            body{background-color: #ddf8ff;}
        </style>

        <title> Go Anagram Webpage </title>
    </head>

    <center> </center><h1> David Cooper's Webpage for Go Anagram Solutions </h1> </center>

    <hr>

    <div class ="myText">
        <h2>Instructions</h2>

        <p>
            Enter one word into the textbox for "Word 1" and the second into the textbox for "Word 2" and click "Submit" to see if the words are anagrams of each other. Note: is not case sensitive.
        </p>

    </div>

    <hr>

    <body>
        <form method="post">
            Word 1:<input type="text" name="ip1">
            Word 2:<input type="text" name="ip2">
            <input type="submit" value="Submit">
        </form>
        Result: {{.Result}}
    </body>

    <p>Github source code: https://github.com/dcooper127/GoPrograms</p>


</html>