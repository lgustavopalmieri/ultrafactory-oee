package location_entity


import "time"

type Location struct {
	Location_ID *int
	Location_Name string 
	Location_Address string 
	Location_Timezone string
	Location_Latitude string
	Location_Longitude string 
	Created_at  time.Time
	Updated_at  time.Time
	Deleted_at  *time.Time
}

