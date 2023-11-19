package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kelvins/geocoder"
	_ "github.com/lib/pq"
)

type CovidJsonRecords []struct {
	Zip_code                           string `json:"zip_code"`
	Week_number                        string `json:"week_number"`
	Week_start                         string `json:"week_start"`
	Week_end                           string `json:"week_end"`
	Cases_weekly                       string `json:"cases_weekly"`
	Cases_cumulative                   string `json:"cases_cumulative"`
	Case_rate_weekly                   string `json:"case_rate_weekly"`
	Case_rate_cumulative               string `json:"case_rate_cumulative"`
	Percent_tested_positive_weekly     string `json:"percent_tested_positive_weekly"`
	Percent_tested_positive_cumulative string `json:"percent_tested_positive_cumulative"`
	Population                         string `json:"population"`
}

type CCVIJsonRecords []struct {
	Geography_type             string `json:"geography_type"`
	Community_area_or_ZIP_code string `json:"community_area_or_zip"`
	Community_name             string `json:"community_area_name"`
	CCVI_score                 string `json:"ccvi_score"`
	CCVI_category              string `json:"ccvi_category"`
}

geocoder.ApiKey = "AIzaSyAf1FsR4foUth1K4kajNTddAJvG5tRGbZE"


