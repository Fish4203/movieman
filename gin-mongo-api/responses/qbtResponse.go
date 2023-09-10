package responses

type QbtResponse struct {
    // basic info 
    Title       string      `json:"name"`
    Size        float64     `json:"size"`
    Catagory    string		`json:"category"` 					
    Progress 	float64		`json:"progress"`
    AddedOn     int 		`json:"added_on"`
    Ratio		float64		`json:"ratio"`
    State 		string		`json:"state"`
    Hash        string      `json:"hash"`
}
