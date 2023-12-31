package responses

type IndexerResponse struct {
    // basic info 
    Title       string                 `json:"title"                    validate:"required"`
    Size        float64                `json:"size"                     validate:"required"`
    Info        string                 `json:"infourl"`
    Date        string                 `json:"publishDate"`
    ReleaseYear int                    `json:"releaseYear"` 
    Catagory    string                 `json:"catagory"                 validate:"required"`


    // torrentinfo 
    Seeders     float64                `json:"seeders"`
    Leachers    float64                `json:"leechers"`
    Indexer     string                 `json:"indexer"`
    Encoding    string                 `json:"encoding"`
    Resolution  string                 `json:"resolution"`
    // season info 
    SeasonNum   string                 `json:"seasonNum"`
    EpisodeNum  string                 `json:"episodeNum"`
    // download 
    Magnet      string                 `json:"magnetUrl,omitempty"`
    Download    string                 `json:"downloadUrl,omitempty"`
}
