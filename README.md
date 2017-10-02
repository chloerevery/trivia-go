# trivia-go

**Get a random trivia question**

curl -s -X GET 'localhost:8000/trivium'

**Add a new trivia question**

curl -X POST -d '{"prompt":"How many books are in the Harry Potter series?","answer":"7"}' localhost:8000/trivium
POST data:

-`prompt` (required)

-`answer` (required)

-`answer_details` (optional)

-`attribution` (optional)
