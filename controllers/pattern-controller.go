package controllers

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/btangkas_agen_frontend/entities"
)

type Responsepattern struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Listpoint   interface{} `json:"listpoint"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Totallose   int         `json:"totallose"`
	Totalwin    int         `json:"totalwin"`
	Time        string      `json:"time"`
}
type Responselistpatterndetail struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Record    interface{} `json:"record"`
	Totallose int         `json:"totallose"`
	Totalwin  int         `json:"totalwin"`
	Time      string      `json:"time"`
}
type Responsepagging struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Time        string      `json:"time"`
}

func Patternhome(c *fiber.Ctx) error {
	type payload_patternhome struct {
		Pattern_search        string `json:"pattern_search"`
		Pattern_search_status string `json:"pattern_search_status"`
		Pattern_page          int    `json:"pattern_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_patternhome)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(Responsepattern{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":       hostname,
			"pattern_search":        client.Pattern_search,
			"pattern_search_status": client.Pattern_search_status,
			"pattern_page":          client.Pattern_page,
		}).
		Post(PATH + "api/pattern")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*Responsepattern)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":      result.Status,
			"perpage":     result.Perpage,
			"totalrecord": result.Totalrecord,
			"totallose":   result.Totallose,
			"totalwin":    result.Totalwin,
			"message":     result.Message,
			"record":      result.Record,
			"listpoint":   result.Listpoint,
			"time":        time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Patternbycode(c *fiber.Ctx) error {
	type payload_patternbycode struct {
		Pattern_code string `json:"pattern_code"`
		Pattern_page int    `json:"pattern_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_patternbycode)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(Responsepagging{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"pattern_code":    client.Pattern_code,
			"pattern_page":    client.Pattern_page,
		}).
		Post(PATH + "api/patternbycode")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*Responsepagging)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":      result.Status,
			"perpage":     result.Perpage,
			"totalrecord": result.Totalrecord,
			"message":     result.Message,
			"record":      result.Record,
			"time":        time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PatternSave(c *fiber.Ctx) error {
	type payload_patternsave struct {
		Page                  string          `json:"page"`
		Sdata                 string          `json:"sdata" `
		Pattern_search        string          `json:"pattern_search" `
		Pattern_page          int             `json:"pattern_page" `
		Pattern_id            string          `json:"pattern_id" `
		Pattern_codepoin      string          `json:"pattern_codepoin" `
		Pattern_status        string          `json:"pattern_status" `
		Pattern_resultcardwin string          `json:"pattern_resultcardwin" `
		Data                  json.RawMessage `json:"data"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_patternsave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	log.Println("Array: ", string(client.Data))
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":       hostname,
			"page":                  client.Page,
			"sdata":                 client.Sdata,
			"pattern_search":        client.Pattern_search,
			"pattern_page":          client.Pattern_page,
			"pattern_id":            client.Pattern_id,
			"pattern_codepoin":      client.Pattern_codepoin,
			"pattern_status":        client.Pattern_status,
			"pattern_resultcardwin": client.Pattern_resultcardwin,
			"pattern_list":          string(client.Data),
		}).
		Post(PATH + "api/patternsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PatternSaveManual(c *fiber.Ctx) error {
	type payload_patternsavemanual struct {
		Page                  string `json:"page"`
		Sdata                 string `json:"sdata" `
		Pattern_search        string `json:"pattern_search" `
		Pattern_page          int    `json:"pattern_page" `
		Pattern_id            string `json:"pattern_id" `
		Pattern_idcard        string `json:"pattern_idcard" `
		Pattern_codepoin      string `json:"pattern_codepoin" `
		Pattern_resultcardwin string `json:"pattern_resultcardwin" `
		Pattern_status        string `json:"pattern_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_patternsavemanual)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":       hostname,
			"page":                  client.Page,
			"sdata":                 client.Sdata,
			"pattern_search":        client.Pattern_search,
			"pattern_page":          client.Pattern_page,
			"pattern_id":            client.Pattern_id,
			"pattern_idcard":        client.Pattern_idcard,
			"pattern_codepoin":      client.Pattern_codepoin,
			"pattern_resultcardwin": client.Pattern_resultcardwin,
			"pattern_status":        client.Pattern_status,
		}).
		Post(PATH + "api/patternsavemanual")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Listpatternhome(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(entities.Home)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
		}).
		Post(PATH + "api/listpattern")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ListpatternSave(c *fiber.Ctx) error {
	type payload_istpatternSave struct {
		Page                      string `json:"page"`
		Sdata                     string `json:"sdata" `
		Listpattern_search        string `json:"listpattern_search" `
		Listpattern_search_status string `json:"listpattern_search_status" `
		Listpattern_page          int    `json:"listpattern_page" `
		Listpattern_id            string `json:"listpattern_id" `
		Listpattern_nmlistpattern string `json:"listpattern_nmlistpattern" `
		Listpattern_status        string `json:"listpattern_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_istpatternSave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":           hostname,
			"page":                      client.Page,
			"sdata":                     client.Sdata,
			"listpattern_search":        client.Listpattern_search,
			"listpattern_search_status": client.Listpattern_search_status,
			"listpattern_page":          client.Listpattern_page,
			"listpattern_id":            client.Listpattern_id,
			"listpattern_nmlistpattern": client.Listpattern_nmlistpattern,
			"listpattern_status":        client.Listpattern_status,
		}).
		Post(PATH + "api/listpatternsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Listpatterndetailhome(c *fiber.Ctx) error {
	type payload_patterndetail struct {
		Listpatterndetail_idlistpattern string `json:"listpatterndetail_idlistpattern" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_patterndetail)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(Responselistpatterndetail{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"listpatterndetail_idlistpattern": client.Listpatterndetail_idlistpattern,
		}).
		Post(PATH + "api/listpatterndetail")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*Responselistpatterndetail)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":    result.Status,
			"message":   result.Message,
			"record":    result.Record,
			"totallose": result.Totallose,
			"totalwin":  result.Totalwin,
			"time":      time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ListpatterndetailSave(c *fiber.Ctx) error {
	type payload_listpatterndetailSave struct {
		Page                            string `json:"page"`
		Sdata                           string `json:"sdata" `
		Listpatterndetail_idlistpattern string `json:"listpatterndetail_idlistpattern" `
		Listpatterndetail_status        string `json:"listpatterndetail_status" `
		Listpatterndetail_idpoin        int    `json:"listpatterndetail_idpoin" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_listpatterndetailSave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":                 hostname,
			"page":                            client.Page,
			"sdata":                           client.Sdata,
			"listpatterndetail_idlistpattern": client.Listpatterndetail_idlistpattern,
			"listpatterndetail_status":        client.Listpatterndetail_status,
			"listpatterndetail_idpoin":        client.Listpatterndetail_idpoin,
		}).
		Post(PATH + "api/listpatterndetailsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ListpatterndetailDelete(c *fiber.Ctx) error {
	type payload_listpatterndetailSave struct {
		Page                            string `json:"page"`
		Sdata                           string `json:"sdata" `
		Listpatterndetail_id            int    `json:"listpatterndetail_id" `
		Listpatterndetail_idlistpattern string `json:"listpatterndetail_idlistpattern" `
		Listpatterndetail_tipe          string `json:"listpatterndetail_tipe" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_listpatterndetailSave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":                 hostname,
			"page":                            client.Page,
			"sdata":                           client.Sdata,
			"listpatterndetail_id":            client.Listpatterndetail_id,
			"listpatterndetail_idlistpattern": client.Listpatterndetail_idlistpattern,
			"listpatterndetail_tipe":          client.Listpatterndetail_tipe,
		}).
		Post(PATH + "api/listpatterndetaildelete")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
