# Getting started
Set the environment variable GOOGLE_CLOUD_PROJECT:
```bash
export GOOGLE_CLOUD_PROJECT=prj-i-calm-tortoise-a9f7
export GOOGLE_CLOUD_QUOTA_PROJECT=$GOOGLE_CLOUD_PROJECT
```
This variable is used by the SDK and also in the `gcloud` commands below.

Clear any previous Application Default Credentials:
```bash
gcloud auth application-default revoke
```

Authenticate and set up your ADC:
```bash
gcloud auth application-default login --scopes=https://www.googleapis.com/auth/cloud-platform,https://www.googleapis.com/auth/spreadsheets.readonly --project=$GOOGLE_CLOUD_PROJECT
```

Set your quota project:
```bash
gcloud config set billing/quota_project $GOOGLE_CLOUD_PROJECT
gcloud auth application-default set-quota-project $GOOGLE_CLOUD_PROJECT
```

Enable the Service Usage API:
```bash
gcloud services enable serviceusage.googleapis.com --project=$GOOGLE_CLOUD_PROJECT
```
This doesn't always work.  If necessary visit this URL https://console.developers.google.com/apis/api/serviceusage.googleapis.com/overview?project=$GOOGLE_CLOUD_PROJECT

Enable the Sheets API:
```bash
gcloud services enable sheets.googleapis.com --project=$GOOGLE_CLOUD_PROJECT
```

Use this environment variable to pass the sheet ID to the program (find it from the URL of your sheet, eg https://docs.google.com/spreadsheets/d/1iKuRqAwWtX-vywfS-Uu5jKP2AfNftYCQelMo3hverBw/edit?gid=0#gid=0)
```bash
export GOOGLE_SHEET_ID=1iKuRqAwWtX-vywfS-Uu5jKP2AfNftYCQelMo3hverBw
```

And another variable to pass the range:
```bash
export GOOGLE_SHEET_RANGE='Sheet1!A1:B5'
```


Build the Golang:
```bash
go build main.go
```

Run the code:
```bash
./main
```

Your output should be something like:
```bash
Data:
A1, B1
A2, B2
A3, B3
A4, B4
A5, B5
```

The complete set of commands is:
```bash
export GOOGLE_CLOUD_PROJECT=prj-i-calm-tortoise-a9f7
export GOOGLE_CLOUD_QUOTA_PROJECT=$GOOGLE_CLOUD_PROJECT

gcloud auth application-default revoke
gcloud auth application-default login --scopes=https://www.googleapis.com/auth/cloud-platform,https://www.googleapis.com/auth/spreadsheets.readonly --project=$GOOGLE_CLOUD_PROJECT

gcloud config set billing/quota_project $GOOGLE_CLOUD_PROJECT
gcloud auth application-default set-quota-project $GOOGLE_CLOUD_PROJECT

gcloud services enable serviceusage.googleapis.com --project=$GOOGLE_CLOUD_PROJECT
gcloud services enable sheets.googleapis.com --project=$GOOGLE_CLOUD_PROJECT

export GOOGLE_SHEET_ID=1iKuRqAwWtX-vywfS-Uu5jKP2AfNftYCQelMo3hverBw
export GOOGLE_SHEET_RANGE='Sheet1!A1:B5'

go build main.go
./main
```
