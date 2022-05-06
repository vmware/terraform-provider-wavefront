package wavefront

// parseStrArr parses a raw interface from d *schema.ResourceData that contains an array of strings
func parseStrArr(raw interface{}) []string {
	var arr []string
	if raw != nil && len(raw.([]interface{})) > 0 {
		for _, v := range raw.([]interface{}) {
			arr = append(arr, v.(string))
		}
	}
	return arr
}
