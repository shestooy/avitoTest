package middlewares

//import (
//	"bytes"
//	"compress/gzip"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestGzipCompression(t *testing.T) {
//	err := storage.MStorage.Init()
//	require.NoError(t, err)
//	flags.Restore = false
//	flags.StorageInterval = 5000
//	handler := GzipMiddleware(http.HandlerFunc(handlers.PostMetricsWithJSON))
//
//	srv := httptest.NewServer(handler)
//	defer srv.Close()
//
//	requestBody := `{
//					  "id": "Alloc",
//					  "type": "gauge",
//					  "value": 100
//					}`
//
//	successBody := `{
//					  "id": "Alloc",
//					  "type": "gauge",
//					  "value": 100
//					}`
//
//	t.Run("sends_gzip", func(t *testing.T) {
//		buf := bytes.NewBuffer(nil)
//		zb := gzip.NewWriter(buf)
//		_, err := zb.Write([]byte(requestBody))
//		require.NoError(t, err)
//		err = zb.Close()
//		require.NoError(t, err)
//
//		r := httptest.NewRequest(http.MethodPost, srv.URL+"/update/", buf)
//		r.Header.Set("Content-Type", "application/json")
//		r.Header.Set("Content-Encoding", "gzip")
//		r.RequestURI = ""
//		r.Header.Set("Accept-Encoding", "")
//
//		resp, err := http.DefaultClient.Do(r)
//		require.NoError(t, err)
//		require.Equal(t, http.StatusOK, resp.StatusCode)
//
//		b, err := io.ReadAll(resp.Body)
//		assert.NoError(t, err)
//		err = resp.Body.Close()
//		require.NoError(t, err)
//		require.JSONEq(t, successBody, string(b))
//	})
//
//	t.Run("accepts_gzip", func(t *testing.T) {
//		buf := bytes.NewBufferString(requestBody)
//		r := httptest.NewRequest(http.MethodPost, srv.URL+"/update/", buf)
//		r.Header.Set("Accept-Encoding", "gzip")
//		r.Header.Set("Content-Type", "application/json")
//		r.RequestURI = ""
//		resp, err := http.DefaultClient.Do(r)
//		require.NoError(t, err)
//		require.Equal(t, http.StatusOK, resp.StatusCode)
//
//		zr, err := gzip.NewReader(resp.Body)
//		assert.NoError(t, err)
//		err = resp.Body.Close()
//		require.NoError(t, err)
//		b, err := io.ReadAll(zr)
//		require.NoError(t, err)
//
//		require.JSONEq(t, successBody, string(b))
//	})
//}
//
