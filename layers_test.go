package geoserver

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetshpFiledsName(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	storename := gsCatalog.GetshpFiledsName("hisham.zip")
	assert.Equal(t, storename, "hisham")
}
func TestGetLayers(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	layers, err := gsCatalog.GetLayers("nurc")
	assert.NotNil(t, layers)
	assert.Nil(t, err)
	gsCatalog = GetCatalog("http://localhost:8080/geoserver_dummy/", "admin", "geoserver")
	layers, err = gsCatalog.GetLayers("nurc_dummy")
	assert.Nil(t, layers)
	assert.NotNil(t, err)
}

func TestGetLayer(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	layer, err := gsCatalog.GetLayer("topp", "tasmania_cities")
	assert.NotNil(t, layer)
	assert.Nil(t, err)
	layer, err = gsCatalog.GetLayer("topp_dummy", "tasmania_cities")
	assert.Equal(t, layer, &Layer{})
	assert.NotNil(t, err)
}
func TestUpdateLayer(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	modified, err := gsCatalog.UpdateLayer("topp", "tasmania_cities", Layer{
		Attribution: &Attribution{
			Title: "Test Title",
		},
	})
	assert.True(t, modified)
	assert.Nil(t, err)
	modified, err = gsCatalog.UpdateLayer("topp_dummy", "tasmania_cities", Layer{
		Attribution: &Attribution{
			Title: "Test Title",
		},
	})
	assert.False(t, modified)
	assert.NotNil(t, err)
}
func TestPublishPostgisLayer(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	conn := DatastoreConnection{
		Name:   "postgis_datastore",
		Port:   5432,
		Type:   "postgis",
		Host:   "postgis",
		DBName: "gis",
		DBPass: "golang",
		DBUser: "golang",
	}
	created, err := gsCatalog.CreateDatastore(conn, "topp")
	assert.True(t, created)
	assert.Nil(t, err)
	published, dbErr := gsCatalog.PublishPostgisLayer("topp", "postgis_datastore", "lbldyt", "lbldyt", "lbldyt")
	assert.True(t, published)
	assert.Nil(t, dbErr)
	published, dbErr = gsCatalog.PublishPostgisLayer("topp", "dummy_store_test", "lbldyt", "lbldyt", "lbldyt")
	assert.False(t, published)
	assert.NotNil(t, dbErr)

}
func TestUploadShapeFile(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	zippedShapefile := filepath.Join(gsCatalog.getGoGeoserverPackageDir(), "testdata", "hurricane_tracks.zip")
	uploaded, err := gsCatalog.UploadShapeFile(zippedShapefile, "shapefileWorkspace", "")
	assert.True(t, uploaded)
	assert.Nil(t, err)
	zippedShapefile = filepath.Join(gsCatalog.getGoGeoserverPackageDir(), "testdata", "hurricane_tracks_dummy.zip")
	uploaded, err = gsCatalog.UploadShapeFile(zippedShapefile, "shapefileWorkspace", "")
	assert.False(t, uploaded)
	assert.NotNil(t, err)
	gsCatalog = GetCatalog("http://localhost:8080/geoserver_dummy/", "admin", "geoserver")
	zippedShapefile = filepath.Join(gsCatalog.getGoGeoserverPackageDir(), "testdata", "hurricane_tracks.zip")
	uploaded, err = gsCatalog.UploadShapeFile(zippedShapefile, "shapefileWorkspace", "")
	assert.False(t, uploaded)
	assert.NotNil(t, err)
}
func TestDeleteLayer(t *testing.T) {
	gsCatalog := GetCatalog("http://localhost:8080/geoserver/", "admin", "geoserver")
	deleted, err := gsCatalog.DeleteLayer("sf", "bugsites", true)
	assert.True(t, deleted)
	assert.Nil(t, err)
	deleted, err = gsCatalog.DeleteLayer("sf_dummy", "bugsites", true)
	assert.False(t, deleted)
	assert.NotNil(t, err)
}
func TestGeoserverImplemetLayerService(t *testing.T) {
	gsCatalog := reflect.TypeOf(&GeoServer{})
	LayerServiceType := reflect.TypeOf((*LayerService)(nil)).Elem()
	check := gsCatalog.Implements(LayerServiceType)
	assert.True(t, check)
}
