package addit

import (
	"net/http"

	"github.com/labstack/echo"
)

func HelloRoot(c echo.Context) error {
	var serverver Info
	serverver.Version = "0.1"
	serverver.VersionString = "hello"

	return c.JSON(http.StatusOK, serverver)
}


func PostUrlIndex(c echo.Context) error {
	var response URLResponse
	request := new(URLRequest)

	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	hs := createHash(request.Url)

	index, has, err := SelectContents(hs)
	if !has || err != nil {
		err = InsertContents(request.Url, hs)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		index, _, err = SelectContents(hs)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		index, err = UpdateContents(hs, index.Count+1)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	response.Count = index.Count
	response.Hash = index.Hash
	response.Index = index.Index

	return c.JSON(http.StatusOK, response)
}

func GetUrlIndex(c echo.Context) error {
	hs := c.Param("hash")
	index, _, err := SelectContents(hs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, index)
}

func PostIndex(c echo.Context) error {
	request := new(ChainRequest)

	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = InsertChainIndex(request.Hash, request.Chainid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, request)
}

func GetIndex(c echo.Context) error {
	hs := c.Param("urlhash")
	indexes, err := SelectChainIndex(hs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, indexes)
}
