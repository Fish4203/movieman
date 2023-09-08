package responses

type IndexerResponse struct {
    // basic info 
    Title       string                 `json:"title,omitempty"`
    Size        float64                `json:"size,omitempty"`
    Info        string                 `json:"infourl,omitempty"`
    Date        string                 `json:"publishDate,omitempty"`
    ReleaseYear int                    `json:"releaseYear,omitempty"` 
    Catagory    string                 `json:"catagory,omitempty"`


    // torrentinfo 
    Seeders     float64                `json:"seeders,omitempty"`
    Leachers    float64                `json:"leechers,omitempty"`
    Indexer     string                 `json:"indexer,omitempty"`
    Encoding    string                 `json:"encoding,omitempty"`
    Resolution  string                 `json:"resolution,omitempty"`
    // season info 
    SeasonNum   int                    `json:"seasonNum,omitempty"`
    EpisodeNum  int                    `json:"episodeNum,omitempty"`
    // download 
    Magnet      string                 `json:"magnetUrl,omitempty"`
    Download    string                 `json:"downloadUrl,omitempty"`
}
