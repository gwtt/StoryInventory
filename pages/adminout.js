(function() {
	const response = {

    "type": "page",
    "data": {
      "count": 1,
    },
    "body": [
      {
      "type": "crud",
      "api": "http://127.0.0.1:3000/objects/getallout?page=${count}",
      "syncLocation": false,
      "loadDataOnce": true,
      "columns": [
        {
          "name": "Id",
          "label": "ID"
        },
        {
          "name": "specie",
          "label": "种类",
          "sortable": true
        },
        {
          "name": "name",
          "label": "名称",
          "sortable": true
        },
        {
          "name": "price",
          "label": "价格（元）",
          "align": "left",
          "sortable": true
        },
        {
          "name": "number",
          "label": "数量",
          "sortable": true
        },
        {
          "name": "unit",
          "label": "单位",
        },
        {
          "name": "out-time",
          "label": "出货时间",
          "sortable": true
        },
        {
          "name": "pur-chase-time",
          "label": "进货时间",
          "sortable": true
        },
        {
          "name": "expiration-time",
          "label": "过期时间",
          "sortable": true
        },
      ]
    }]
  }

	window.jsonpCallback && window.jsonpCallback(response);
})();
