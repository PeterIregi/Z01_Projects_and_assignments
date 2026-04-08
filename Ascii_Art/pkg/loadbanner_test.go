package pkg

import (
    "os"
    "testing"
)



func TestLoadBanner_FileNotFound(t *testing.T) {
    filename := "../standard.txt"
    // Test how the function handles a missing file
    t.Run("file does not exist", func(t *testing.T) {
        bannerMap := LoadBanner("non_existent_file.txt")

        if bannerMap != nil {
            t.Errorf("Expected nil return for non-existent file, got %v", bannerMap)
        }
    })
    // Test that the correct number of ascii characters are stored
    t.Run("file has correct length", func(t *testing.T) {
        bannerMap := LoadBanner(filename)

        if len(bannerMap) != 95 {
            t.Errorf("Expected 95, got %v", len(bannerMap))
        }

    })
    // Test for an empty file
    t.Run("empty banner file", func(tq *testing.T) {
        tmpFile := "test_banner.txt"

        file, _ := os.Create(tmpFile)
        file.Close()
        defer os.Remove(tmpFile)

        bannerMap := LoadBanner(tmpFile)
        if bannerMap != nil {
        t.Errorf("Expected bannerMap to be nil, but got %v", bannerMap)
        }
    })
}