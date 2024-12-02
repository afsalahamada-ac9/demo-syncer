package rds_to_sf_entity

type Data struct {
	Object string
	Items  []interface{} // this is a dynamic type, but then we'll limit the parse possibilities
}
