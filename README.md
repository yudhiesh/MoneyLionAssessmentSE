## Test Cases

**Test cases are run under [controller_test.go](https://github.com/yudhiesh/MoneyLionAssessmentSE/blob/b48ed9bb5e99c9fe450b327cbf1857ef8be8ff40/controller/controller_test.go).**

1. TestCanGetAccess tests the `GET` request to `/feature`

```
❯ go test -v -run="TestGetCanAcess" ./controller/
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

2. TestInsertFeature tests the `POST` request to `/feature`

```
❯ go test -v -run="TestInsertFeature" ./controller/
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
