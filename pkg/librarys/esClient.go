package librarys

import (
	"encoding/json"
	"context"
	"github.com/spf13/cast"
	global "github.com/EDDYCJY/go-gin-example/pkg"
)

type EsClientHandler struct {

}

//根据id获取数据
func (this *EsClientHandler) EsClientGetInfoById(id, indexName, typeName string) map[string]interface{} {
	var result = make(map[string]interface{})
	res, err := global.ES.Get().Index(indexName).Type(typeName).Id(id).Do(context.Background())
	if err != nil {
		return result
	}
	if res == nil {
		return result
	}
	data, _ := res.Source.MarshalJSON()
	json.Unmarshal(data, &result)
	return result
}


//插入数据
func (this *EsClientHandler) EsClientInsert(id, indexName, typeName string, doc interface{}) error {
	_, err := global.ES.Index().Index(indexName).Type(typeName).Id(cast.ToString(id)).BodyJson(doc).Refresh("wait_for").Do(context.Background())
	//_, err := global.ES.Update().Index(indexName).Type(typeName).Id(id).Doc(doc).Refresh("wait_for").Do(context.Background())
	return err
}

//根据id更新数据
func (this *EsClientHandler) EsClientUpdateById(id, indexName, typeName string, doc map[string]interface{}) error {
	_, err := global.ES.Update().Index(indexName).Type(typeName).Id(id).Doc(doc).Refresh("wait_for").Do(context.Background())
	return err
}

//根据id删除数据
func (this *EsClientHandler) DeleteById(id, indexName, typeName string) error {
	_, err := global.ES.Delete().Index(indexName).Type(typeName).Id(id).Do(context.Background())
	return err
}

/*

func (this *ShopsEsService) ShopsSearch(keyword string, lat float64, lon float64, page int, scrollId string, isOpen string, isDelivery string, pageSize int, cover string, shopIdExclude int, isScore string, keywordType string, distanceString string, distanceType string) map[string]interface{} {
	var esSearchService *elastic.ScrollService
	//var esSearchServiceSearch *elastic.SearchService
	boolQuery := elastic.NewBoolQuery()

	if len(keywordType) > 0 {
		var keywordQuery = elastic.NewMultiMatchQuery(keyword).Field("shop_name").Field("address").TieBreaker(0.3).MinimumShouldMatch("30%")
		boolQuery.Should(keywordQuery)
	} else {
		var keywordQuery = elastic.NewMatchQuery("name_address", strings.ToLower(keyword))
		boolQuery.Should(keywordQuery)
	}

	if len(distanceType) > 0 {
		distanceQuery := elastic.NewGeoDistanceQuery("location")
		distanceQuery = distanceQuery.Lat(lat)
		distanceQuery = distanceQuery.Lon(lon)
		distanceQuery = distanceQuery.Distance(distanceString)
		//boolQuery = boolQuery.Must(elastic.NewMatchAllQuery())
		boolQuery.Filter(distanceQuery)
	}

	if len(cover) > 0 {
		existsQuery := elastic.NewExistsQuery("cover")
		boolQuery.MustNot(existsQuery)
	}

	if shopIdExclude > 0 {
		shopIdQuery := elastic.NewTermQuery("shop_id", shopIdExclude)
		boolQuery.MustNot(shopIdQuery)
	}

	if len(isOpen) > 0 {
		isOpenQuery := elastic.NewTermQuery("is_open", 1)
		boolQuery.Must(isOpenQuery)
	}


	//boolQuery.Must(subBoolQuery)
	if len(isDelivery) > 0 {
		isDeliveryQuery := elastic.NewTermQuery("is_delivery", 1)
		boolQuery.Must(isDeliveryQuery)
	}

	//boolQuery.Filter(distanceQuery)
	//boolQuery.Must(distanceQuery)
	//searchSource := elastic.NewSearchSource().TrackScores(true)

	//esSearchService = global.ES.Scroll().Index(consts.ShopsEsIndex).Type(consts.ShopsEsType).SearchSource(searchSource).Query(boolQuery).Scroll("60m").Size(pageSize).ScrollId(scrollId)
	esSearchService = global.ES.Scroll().Index(consts.ShopsEsIndex).Type(consts.ShopsEsType).Query(boolQuery).Scroll("60m").Size(pageSize).ScrollId(scrollId)
	if len(isScore) > 0 {
		esSearchService = esSearchService.Sort("shop_id", false)
	} else {
		geoDistanceSorter := elastic.NewGeoDistanceSort("location").Point(lat, lon).Unit("km").Asc()
		esSearchService = esSearchService.SortBy(geoDistanceSorter)
		esSearchService = esSearchService.Sort("shop_score", false)
		esSearchService = esSearchService.Sort("shop_id", true)
	}

	searchResult, err := esSearchService.Do(context.Background())
	if err != nil {
		panic(err)
	}
	data := helpers.StructToMap(searchResult)
	return data
}



func (this *ShopsEsService) ShopsNation(lat float64, lon float64, notIn string, distanceString string ,pageSize int ,trackType string,certifiedType string,termsType string) map[string]interface{} {
	var esSearchService *elastic.ScrollService
	//var esSearchServiceSearch *elastic.SearchService
	boolQuery := elastic.NewBoolQuery()
	distanceQuery := elastic.NewGeoDistanceQuery("location")
	distanceQuery = distanceQuery.Lat(lat)
	distanceQuery = distanceQuery.Lon(lon)
	distanceQuery = distanceQuery.Distance(distanceString)
	//boolQuery = boolQuery.Must(elastic.NewMatchAllQuery())
	boolQuery.Filter(distanceQuery)
	var shopIds = make([]int, 0)
	json.Unmarshal([]byte(notIn), &shopIds)
	termsNotIn := helpers.ToInterfaceSlice(shopIds)
	shopIdQuery := elastic.NewTermsQuery("shop_id",termsNotIn... )
	if(termsType == "mustnot"){
		boolQuery.MustNot(shopIdQuery)
	}

	if(termsType == "must"){
		boolQuery.Must(shopIdQuery)
	}


	if certifiedType == "certified" {
		isOpenQuery := elastic.NewTermQuery("is_certified", 1)
		boolQuery.Must(isOpenQuery)
	}
	if certifiedType == "uncertified" {
		isOpenQuery := elastic.NewTermQuery("is_certified", 0)
		boolQuery.Must(isOpenQuery)
	}



	if len(trackType) > 0 {
		searchSource := elastic.NewSearchSource().TrackScores(true)
		esSearchService = global.ES.Scroll().Index(consts.ShopsEsIndex).Type(consts.ShopsEsType).SearchSource(searchSource).Query(boolQuery).Scroll("60m").Size(pageSize).ScrollId("")
	} else {
		esSearchService = global.ES.Scroll().Index(consts.ShopsEsIndex).Type(consts.ShopsEsType).Query(boolQuery).Scroll("60m").Size(pageSize).ScrollId("")
		esSearchService = esSearchService.Sort("shop_id", true)
	}

	geoDistanceSorter := elastic.NewGeoDistanceSort("location").Point(lat, lon).Unit("km").Asc()
	esSearchService = esSearchService.SortBy(geoDistanceSorter)
	searchResult, err := esSearchService.Do(context.Background())
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	data := helpers.StructToMap(searchResult)
	return data
}
 */