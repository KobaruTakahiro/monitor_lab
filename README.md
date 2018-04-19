# monitor_lab

## SAM
## generate local sam event

```
sam local generate-event schedule > events.json
```

### invoke lambda function
```
 sam local invoke MonitorCollector -e events.json
 ```

## add assume Role

add another aws account.
next,Trust rleationship and edit trust releationship
access iam usr arn