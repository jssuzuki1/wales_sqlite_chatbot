## Github Link: 

https://github.com/jssuzuki1/wales_sqlite_chatbot

## What this app does

Launch the app by navigating to build\bin\go_chatbot.exe

The user is prompted with a text box, asking for a term related to Go. If the term belongs to our existing db, then it qualifies.

The user may access this list of terms by navigating to 'QandA.csv', where the first column represents the domain of available terms that will return a result in the application.
s
## How it works

There is a database called 'QandA.db' located in the 'tutorDB' folder. 

The main.go program in the root directory generates this database using the QandA.csv file, also located in the root directory. This CSV file originated from https://github.com/ThomasWMiller/jump-start-sqlite. Both the 'question' and 'answer' columns generated in the database strip all quotation marks so that the user does not have to type their answers in quotation marks.

Querying the database occurs in app.go. The 'query' variable in the 'Search' function is input by the user. The search function returns the term the user searched for and the definition of that term.

## Further Development of the Application
 
One option is to expand the data set. Webscrape the pages off of go for terms and their definitions and add them to our CSV files (or change the current file to JSON). Ideally, this would result in a database of hundreds, if not thousands, of searchable terms.

Another option would be to ping a ChatGPT API for answers to more complex questions that we cannot anticipate. For example, if someone asks a hyperspecific question about whether their code is functional, we are unlikely to anticipate their particular case for an answer. In this method, SQLite, a local database, is not suitable for the job.

Considering that Go is extremely fast and versatile, we should stick to using Go. But Python has a lot of useful packages for preprocessing language data that Go might not have. That is a consideration.

But I will take a more categorical tone since I'm in a Go class right now. : vive l'Go! 

## Image source:

https://knowyourmeme.com/memes/dramatic-chipmunk

It's not a gopher, but close enough. 