(function() {
	const response = {
    
    "type": "page",
    "data": {
      "count": 1,
    },
    "body": [
      {    
        "label": "新增",
        "type": "button",
        "actionType": "dialog",
        "level": "primary",
        "className": "m-b-sm",
        "dialog": {
          "title": "新增表单",
          "body": {
            "type": "form",
            "api": "post:http://127.0.0.1:3000/objects/create",
            "body": [
              {
                "type": "select",
                "name": "specie",
                "label": "种类",
                "options": [
                  {
                    "label": "水果",
                    "value": "水果"
                  },
                  {
                    "label": "日用品",
                    "value": "日用品"
                  },
                  {
                    "label": "电子产品",
                    "value": "电子产品"
                  },
                  {
                    "label": "零食",
                    "value": "零食"
                  },                  
                  {
                    "label": "玩具",
                    "value": "玩具"
                  },
                  {
                    "label": "奢侈品",
                    "value": "奢侈品"
                  },
                  {
                    "label": "学习用品",
                    "value": "学习用品"
                  },
                  {
                    "label": "工料品",
                    "value": "工料品"
                  },
                ]
              },
              {
                "type": "input-text",
                "name": "name",
                "label": "名称"
              },
              {
                "type": "input-number",
                "name": "price",
                "label": "价格"
              },
              {
                "type": "input-number",
                "name": "number",
                "label": "数量"
              },
              {
                "type": "select",
                "name": "unit",
                "label": "单位",
                "options": [
                  {
                    "label": "克",
                    "value": "克",
                  },
                  {
                    "label": "千克",
                    "value": "千克",
                  },
                  {
                    "label": "斤",
                    "value": "斤",
                  },
                  {
                    "label": "两",
                    "value": "两",
                  },
                  {
                    "label": "个",
                    "value": "个",
                  },
                  {
                    "label": "瓶",
                    "value": "瓶",
                  },
                  {
                    "label": "辆",
                    "value": "辆",
                  },
                  {
                    "label": "吨",
                    "value": "吨",
                  },
                ]
              },
              {
                "type": "input-date",
                "name": "pur-chase-time",
                "label": "进货时间",
                "format": "YYYY-MM-DD"
              },
              {
                "type": "input-date",
                "name": "expiration-time",
                "label": "过期时间",
                "format": "YYYY-MM-DD"
              }
            ]
          }
        }
      },
      {
      "type": "crud",
      "api": "http://127.0.0.1:3000/objects/getall",
      "syncLocation": false,
      "headerToolbar": [
        {
          "type": "button",
          "actionType": "ajax",
          "label": "过期邮箱警报",
          "api": "Get:http://127.0.0.1:3000/users/alert"
        }
      ],
      "loadDataOnce": true,
      "columns": [
        {
          "name": "Id",
          "label": "ID",

        },
        {
          "name": "specie",
          "label": "种类",
          "sortable": true,
          "searchable": true
        },
        {
          "name": "name",
          "label": "名称",
          "sortable": true,
          "searchable": true
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
          "name": "pur-chase-time",
          "label": "进货时间",
          "sortable": true,
          "searchable": true
        },
        {
          "name": "expiration-time",
          "label": "过期时间",
          "sortable": true,
          "searchable": true
        },
        {
          "type": "operation",
          "label": "操作",
          "buttons": [
            {
              "label": "修改",
              "type": "button",
              "actionType": "drawer",
              "drawer": {
                "title": "进货",
                "body": {
                  "type": "form",
                  "api": "post:http://127.0.0.1:3000/objects/update?id=${Id}",
                  "body": [
                    {
                      "type": "select",
                      "name": "specie",
                      "label": "种类",
                      "options": [
                        {
                          "label": "水果",
                          "value": "水果"
                        },
                        {
                          "label": "日用品",
                          "value": "日用品"
                        },
                        {
                          "label": "电子产品",
                          "value": "电子产品"
                        },
                        {
                          "label": "零食",
                          "value": "零食"
                        },                  
                        {
                          "label": "玩具",
                          "value": "玩具"
                        },
                        {
                          "label": "奢侈品",
                          "value": "奢侈品"
                        },
                        {
                          "label": "学习用品",
                          "value": "学习用品"
                        },
                      ]
                    },
                    {
                      "type": "input-text",
                      "name": "name",
                      "label": "名称"
                    },
                    {
                      "type": "input-number",
                      "name": "price",
                      "label": "价格"
                    },
                    {
                      "type": "input-number",
                      "name": "number",
                      "label": "数量"
                    },
                    {
                      "type": "select",
                      "name": "unit",
                      "label": "单位",
                      "options": [
                        {
                          "label": "克",
                          "value": "克",
                        },
                        {
                          "label": "千克",
                          "value": "千克",
                        },
                        {
                          "label": "斤",
                          "value": "斤",
                        },
                        {
                          "label": "两",
                          "value": "两",
                        },
                        {
                          "label": "个",
                          "value": "个",
                        },
                        {
                          "label": "瓶",
                          "value": "瓶",
                        },
                        {
                          "label": "辆",
                          "value": "辆",
                        },
                      ]
                    },
                    {
                      "type": "input-date",
                      "name": "pur-chase-time",
                      "label": "进货时间",
                      "format": "YYYY-MM-DD"
                    },
                    {
                      "type": "input-date",
                      "name": "expiration-time",
                      "label": "过期时间",
                      "format": "YYYY-MM-DD"
                    }
                  ]
                }
              }
            },
            {
              "label": "删除",
              "type": "button",
              "actionType": "ajax",
              "level": "danger",
              "confirmText": "确认要删除？",
              "api": "delete:http://127.0.0.1:3000/objects/delete?id=${Id}"
            },
            {
              "label": "出货",
              "type": "button",
              "actionType": "ajax",
              "level": "primary",
              "confirmText": "确认要出货？",
              "api": "post:http://127.0.0.1:3000/objects/out?id=${Id}"
            }
          ]
        },
      ]
    }]
  }

	window.jsonpCallback && window.jsonpCallback(response);
})();
