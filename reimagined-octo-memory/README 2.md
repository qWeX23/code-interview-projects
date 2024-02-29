# * Applicant Take Home

## Running the app
`docker compose up --build`

## Introduction

The * Applicant Take-Home is a programming exercise that is built to
allow you to display your programming abilities. At * we primarily use
Kotlin, Java, and Go. We would prefer you use one of these languages
for this task. However, we ask that you please use the language that you are
most comfortable with for this assessment.

## Time Limit

4 Hours.

Please don't take any longer as we want to be respectful of your time.

Please tell us how long you spent on this project if less than four hours (see Reflection below).

## Reflection

The client would like to know what you would do on this project if
given more time as well as what you thought the most and least
challenging parts of the project were. Please write some thoughts on
both of these topics in a file called `REFLECTION.md` (bulleted lists
are okay).

## Data

Included is a file named `MOCK_DATA.json` (sample record below).  Use this data to seed a data store that backs your application (requirements below).

```json
{
    "id": 2,
    "first_name": "Thacher",
    "last_name": "Linnett",
    "gender": "Male",
    "phone_number": "194-129-0724",
    "email": "tlinnett1@upenn.edu",
    "address": "278 Talmadge Park",
    "visit_date": "9/13/2017",
    "diagnosis": "A968",
    "drug_code": "43742-0299",
    "additional_information": [
        {
            "notes": "In hac habitasse platea dictumst. Etiam faucibus cursus urna. Ut tellus.\n\nNulla ut erat id mauris vulputate elementum. Nullam varius. Nulla facilisi.\n\nCras non velit nec nisi vulputate nonummy. Maecenas tincidunt lacus at velit. Vivamus vel nulla eget eros elementum pellentesque.",
            "new_patient": false,
            "race": "Puerto Rican",
            "ssn": "735-91-0685"
        }
    ]
}
```

## Your mission:

A hospital wants a RESTful API interface into their patient records. 

There is no need for a user interface; just the API.

You may utilize any HTTP Client that you desire to test your API (eg. POSTman, curl, etc.).

As mentioned previously, you can use the provided seed data in `MOCK_DATA.json` to seed your datastore.  There is a JSON array with 1000 patient records within the file.

(Often, the patients don't always supply the information you request.  Sometimes
their patient record is incomplete, so make sure you handle empty fields! However,
the diagnosis is a required field. Any patients missing a diagnosis should log
a rejection message and should not be stored.  The API should return an appropriate validation message stating that the required field is missing).

The client should be able to perform both searching and CRUD operations as
defined below.

- Your API server may run locally.  There is no need to host this anywhere.  
- If your server requires special instructions to run, please provide them.
- Your server should run on port `3000`.


### Searching

The client needs to be able to search for patients by any of their attributes.
The client will run a command similar to the examples below and should get a
list of patient `id` numbers in return.

```bash
curl -G -v "http://localhost:3000/search" --data-urlencode "first_name=Rollo"
curl -G -v "http://localhost:3000/search" --data-urlencode "gender=Male"
curl -G -v "http://localhost:3000/search" --data-urlencode "last_name=Methuen"
```

### CRUD Operations

The client also needs to be able to perform CRUD operations on the patients in
their database. Please see the details of what they need to accomplish for each
method below. The client should always get back one of the three following
responses:

- json object of the new/updated patient information
- empty json object (`{}`) if the patient information was just deleted
- json object with an error if the patient doesn't exist or validation failure (e.x. `{"error": "Patient Does Not Exist"}`)
  
**Create**
  (_Hint: Remember to account for partial uploads of new patient information like
  the example below._)

```bash
curl -d '{"first_name":"Bob", "last_name":"Saget"}' -H "Content-Type: application/json" -X POST http://localhost:3000/patient
```

**Read**

```bash
curl http://localhost:3000/patient/1
```

**Update**

```bash
curl -d '{"name":"Bob"}' -H "Content-Type: application/json" -X PUT http://localhost:3000/patient/1
```

**Delete**

```bash
curl -X DELETE http://localhost:3000/patient/1
```

## Reminders!

- Don't forget to `git push` regularly.
- Have fun!
- Contact Ben Bowman (ben@*.com) if you experience any issues.

