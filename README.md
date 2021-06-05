# MoneyLion AI Software Engineer Assessment

As Product Manager, I would like to manage users’ accesses to new features via feature switches,i.e. enabling/disabling certain feature based on a user’s email and feature names).

## API Reference

#### Get the boolean value of whether a user has access to a particular feature

```http
  GET https://api-moneylion.herokuapp.com/feature
```

| Parameter     | Type     | Description                       |
| :------------ | :------- | :-------------------------------- |
| `email`       | `string` | **Required**. Email of the user   |
| `featureName` | `string` | **Required**. Name of the feature |

#### Update the access a user has to a particular feature

```http
  POST https://api-moneylion.herokuapp.com/feature
```

| Parameter     | Type      | Description                                                                |
| :------------ | :-------- | :------------------------------------------------------------------------- |
| `featureName` | `string`  | **Required**. Name of the feature                                          |
| `email`       | `string`  | **Required**. Email of the user                                            |
| `can_access`  | `boolean` | **Required**. Change the access a user has to a feature this boolean value |

## Usage/Examples

Example 1:

**GET** request to `https://api-moneylion.herokuapp.com/feature?email=test3@gmail.com&featureName=automated-investing`

Result:

```
{
    "can_access":false
}
```

Example 2:

**GET** request to `https://api-moneylion.herokuapp.com/feature`

Result:

```
{
    "error":"Missing URL query parameters email/featureName"
}
```

Example 3:

**POST** request to `https://api-moneylion.herokuapp.com/feature`

Request body:

```
{
    "featureName": "automated-investing",
    "email": "test1@gmail.com",
    "can_access": true
}
```

Result:

```
Status: 304 Not Modified
```

Example 4:

**POST** request to `https://api-moneylion.herokuapp.com/feature`

Request body:

```
{
    "featureName": "automated-investing",
    "email": "test1@gmail.com",
    "can_access": false
}
```

Result:

```
Status: 200 OK
```

Example 5:

**POST** request to `https://api-moneylion.herokuapp.com/feature`

Request body:

```
{}
```

Result:

```
Status: 304 Not Modified
```

Example 6:

**POST** request to `https://api-moneylion.herokuapp.com/feature`

Request body:

```
{
    "featureName": "automated-investing",
    "email": "test1@gmail.com"
}
```

Result:

```
Status: 304 Not Modified
```

## Database

To simulate the scenario I have created two tables called _users_ and _features_ where _users_ contains the _id_ of the user and the _email_.

The _features_ table contains all the features and whether or not a user has access to those features.

[![Screenshot-2021-06-05-at-10-09-10-PM.png](https://i.postimg.cc/rmcM85xV/Screenshot-2021-06-05-at-10-09-10-PM.png)](https://postimg.cc/PL6BypS9)

- users

```
+----+-----------------+
| id | email           |
+----+-----------------+
|  1 | test1@gmail.com |
|  2 | test2@gmail.com |
|  3 | test3@gmail.com |
|  4 | test4@gmail.com |
|  5 | test5@gmail.com |
+----+-----------------+
```

- features

```
+------------+---------+---------------------+------------+
| feature_id | user_id | feature_name        | can_access |
+------------+---------+---------------------+------------+
|          1 |       1 | automated-investing |          1 |
|          2 |       1 | crypto              |          0 |
|          3 |       2 | crypto              |          0 |
|          4 |       3 | automated-investing |          0 |
|          5 |       4 | automated-investing |          1 |
|          7 |       1 | financial-tracking  |          1 |
|          8 |       2 | financial-tracking  |          0 |
|          9 |       3 | financial-tracking  |          1 |
|         10 |       4 | financial-tracking  |          0 |
+------------+---------+---------------------+------------+
```

## Test Cases

**Test cases are run under [controller_test.go](https://github.com/yudhiesh/MoneyLionAssessmentSE/blob/b48ed9bb5e99c9fe450b327cbf1857ef8be8ff40/controller/controller_test.go).**

1. **TestCanGetAccess** tests the `GET` request to `/feature`

```
> go test -v -run="TestGetCanAcess" ./controller/
=== RUN   TestGetCanAcess
=== RUN   TestGetCanAcess/Missing_Entire_URL_Parameter
=== RUN   TestGetCanAcess/Missing_Email_URL_Parameter
=== RUN   TestGetCanAcess/Missing_featureName_URL_Parameter
=== RUN   TestGetCanAcess/Email_does_not_exist
=== RUN   TestGetCanAcess/Email_and_FeatureName_exist
=== RUN   TestGetCanAcess/Email_and_FeatureName_exist#01
=== RUN   TestGetCanAcess/FeatureName_does_not_exist
--- PASS: TestGetCanAcess (0.01s)
    --- PASS: TestGetCanAcess/Missing_Entire_URL_Parameter (0.00s)
    --- PASS: TestGetCanAcess/Missing_Email_URL_Parameter (0.00s)
    --- PASS: TestGetCanAcess/Missing_featureName_URL_Parameter (0.00s)
    --- PASS: TestGetCanAcess/Email_does_not_exist (0.00s)
    --- PASS: TestGetCanAcess/Email_and_FeatureName_exist (0.00s)
    --- PASS: TestGetCanAcess/Email_and_FeatureName_exist#01 (0.00s)
    --- PASS: TestGetCanAcess/FeatureName_does_not_exist (0.00s)
PASS
ok  	github.com/yudhiesh/api/controller	0.150s
```

2. **TestInsertFeature** tests the `POST` request to `/feature`

```
> go test -v -run="TestInsertFeature" ./controller/
=== RUN   TestInsertFeature
=== RUN   TestInsertFeature/Missing_entire_body
=== RUN   TestInsertFeature/Missing_body_1
=== RUN   TestInsertFeature/Missing_body_2
=== RUN   TestInsertFeature/Missing_body_3
=== RUN   TestInsertFeature/Incorrect_value_for_parameters_1
=== RUN   TestInsertFeature/Incorrect_user_email
=== RUN   TestInsertFeature/Incorrect_can_access
=== RUN   TestInsertFeature/Correct_case_1
=== RUN   TestInsertFeature/Correct_case_2
--- PASS: TestInsertFeature (0.05s)
    --- PASS: TestInsertFeature/Missing_entire_body (0.00s)
    --- PASS: TestInsertFeature/Missing_body_1 (0.00s)
    --- PASS: TestInsertFeature/Missing_body_2 (0.00s)
    --- PASS: TestInsertFeature/Missing_body_3 (0.00s)
    --- PASS: TestInsertFeature/Incorrect_value_for_parameters_1 (0.00s)
    --- PASS: TestInsertFeature/Incorrect_user_email (0.00s)
    --- PASS: TestInsertFeature/Incorrect_can_access (0.00s)
    --- PASS: TestInsertFeature/Correct_case_1 (0.02s)
    --- PASS: TestInsertFeature/Correct_case_2 (0.01s)
PASS
ok  	github.com/yudhiesh/api/controller	0.577s
```

## Tech Stack

**Client:** Golang

**Database:** MySQL

**Deployment:** Heroku
