package a_test

import (
	project2 "bws_microservice_url/main/src/core/urls"
	"bws_microservice_url/main/src/entity"
	"testing"
	"time"

	"github.com/bindways/bw_microservice_share/bw_helper/bw_time_helper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBwsProjectService_InsertKeywordTrack(t1 *testing.T) {
	idProject, _ := primitive.ObjectIDFromHex("64e82827ef2b3c9dfd72c9f8")
	keywordTrack := entity.BwsKeywordTrack{
		Keyword:     "bindways partner",
		TrackAtDate: bw_time_helper.TimePlusXMonths(1),
	}
	siteTrack := entity.BwsUrlTrack{
		CurrentPosition: 12,
		LastPosition:    0,
		Url:             "/teste/url",
		CreateAtDate:    time.Now(),
	}
	service := new(project2.BwsUrlTrackService).Constructor1()
	project, err := service.InsertKeywordTrack(idProject, keywordTrack, siteTrack)
	assert.Nil(t1, err)
	assert.EqualValues(t1, idProject, project.Id)
	assert.EqualValues(t1, keywordTrack.Keyword, project.KeywordTracks[0].Keyword)
	assert.EqualValues(t1, keywordTrack.TrackAtDate, project.KeywordTracks[0].TrackAtDate)
	assert.EqualValues(t1, siteTrack.Url, project.KeywordTracks[0].SiteTracks[0].Url)
	assert.EqualValues(t1, siteTrack.CurrentPosition, project.KeywordTracks[0].SiteTracks[0].CurrentPosition)
}

func TestBwsProjectService_UpdateKeywordTrack(t1 *testing.T) {
	idProject, _ := primitive.ObjectIDFromHex("64e82827ef2b3c9dfd72c9f8")
	keywordTrack := entity.BwsKeywordTrack{
		Keyword:     "bindways partner",
		TrackAtDate: bw_time_helper.TimePlusXMonths(1),
	}
	siteTrack := entity.BwsUrlTrack{
		CurrentPosition: 232,
		LastPosition:    0,
		Url:             "/teste/url",
		CreateAtDate:    time.Now(),
	}
	service := new(project2.BwsUrlTrackService).Constructor1()
	project, err := service.UpdateKeywordTrack(idProject, keywordTrack, siteTrack)
	assert.Nil(t1, err)
	assert.EqualValues(t1, idProject, project.Id)
	assert.EqualValues(t1, keywordTrack.Keyword, project.KeywordTracks[0].Keyword)
	assert.EqualValues(t1, keywordTrack.TrackAtDate, project.KeywordTracks[0].TrackAtDate)
	assert.EqualValues(t1, siteTrack.Url, project.KeywordTracks[0].SiteTracks[0].Url)
	assert.EqualValues(t1, siteTrack.CurrentPosition, project.KeywordTracks[0].SiteTracks[0].CurrentPosition)
}

func TestBwsProjectService_UpdateTrackDate(t1 *testing.T) {
	idProject, _ := primitive.ObjectIDFromHex("64e82827ef2b3c9dfd72c9f8")
	service := new(project2.BwsUrlTrackService).Constructor1()
	_, err := service.UpdateTrackDate(idProject, "bindways partner", bw_time_helper.PlusXDays(time.Now(), 2))
	assert.Nil(t1, err)
}
