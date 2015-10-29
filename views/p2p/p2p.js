var $select_titles = new Array();

var $wealth_contents = new Array();

(function () {
    $.ajax({
        type: 'GET',
        url: "http://localhost:7070/v1/lov/types",
        dataType: "json",
        async: false,
        contentType: "application/json;charset=UTF-8",
        success: function (data) {
            $data = $(data)
            for (var i = 0; i < $data.length; i++) {
                var obj = new Object();

                obj.child_title = $data[i].Name;

                var child_content = new Array();
                child_content[0]="不限"
                $.ajax({
                    type: "GET",
                    url: "http://localhost:7070/v1/lov/list/" + $data[i].Code,
                    dataType: "json",
                    async: false,
                    success: function (content) {
                        $content = $(content);
                        for (var j = 0; j < $content.length; j++) {
                            child_content[j+1] = $content[j].Name;
                        }
                        obj.child_content = child_content;
                    }

                });

                obj.ifdisplay = false;
                obj.index = i;
                obj.chosen = 0;
                $select_titles[i] = obj;
            }
        }

    });


    $.ajax({
        type: 'GET',
        url: "http://localhost:7070/v1/invest/list",
        dataType: "json",
        async: false,
        contentType: "application/json;charset=UTF-8",
        success: function (data) {
            $data = $(data)
            for (var i = 0; i < $data.length; i++) {
                var obj = new Object();

                obj.title = $data[i].Name;
                obj.url = $data[i].Master.Official_url;
                obj.profit = $data[i].Rate;
                obj.money = $data[i].Amount_min;
                obj.time = $data[i].Duration_min;
                obj.ifdisplay = "true"

                $wealth_contents[i] = obj;
            }
        }

    });

})()

var p2p_filter = new Vue({
    el: '#p2p_filter',
    data: {
        state: [false, false, false],
        select_titles: $select_titles
    },
    methods: {
        toggle: function (item) {
            if (item.ifdisplay) {
                item.ifdisplay = !item.ifdisplay;
            }
            else {
                p2p_filter.select_titles[0].ifdisplay = false;
                p2p_filter.select_titles[1].ifdisplay = false;
                p2p_filter.select_titles[2].ifdisplay = false;
                item.ifdisplay = true;
            }
        },
        cancel: function () {
            p2p_filter.select_titles[0].ifdisplay = false;
            p2p_filter.select_titles[1].ifdisplay = false;
            p2p_filter.select_titles[2].ifdisplay = false;
        },
        select: function (selecttype, filtertype, filtervalue, item) {
            if (selecttype == 0) {
                item.child_title = item.child_title_bk;
                p2p_filter.select_titles[filtertype].chosen = selecttype;
            }
            else {
                item.child_title = filtervalue;
                p2p_filter.select_titles[filtertype].chosen = selecttype;
            }
            p2p_filter.select_titles[0].ifdisplay = false;
            p2p_filter.select_titles[1].ifdisplay = false;
            p2p_filter.select_titles[2].ifdisplay = false;
            // console.log(p2p_filter.select_titles[0].chosen);
            // console.log(p2p_filter.select_titles[1].chosen);
            // console.log(p2p_filter.select_titles[2].chosen);
            set_filter(p2p_filter.select_titles[0].chosen, p2p_filter.select_titles[1].chosen, p2p_filter.select_titles[2].chosen);
        }
    },
})
var wealth_result = new Vue({
    el: '#wealth_result',
    data: {
        wealth_contents: $wealth_contents
    }
})

var set_filter = function (money, time, profit) {
    for (item in wealth_result.wealth_contents) {
        flag = false;
        switch (money) {
            case 0:
                flag = true;
                break;
            case 1:
                if (wealth_result.wealth_contents[item].money <= 50000) flag = true;
                break;
            case 2:
                if (wealth_result.wealth_contents[item].money <= 100000 && wealth_result.wealth_contents[item].money > 50000) flag = true;
                break;
            case 3:
                if (wealth_result.wealth_contents[item].money <= 200000 && wealth_result.wealth_contents[item].money > 100000) flag = true;
                break;
            case 4:
                if (wealth_result.wealth_contents[item].money <= 500000 && wealth_result.wealth_contents[item].money > 200000) flag = true;
                break;
            case 5:
                if (wealth_result.wealth_contents[item].money > 500000) flag = true;
                break;
        }
        if (!flag) {
            wealth_result.wealth_contents[item].ifdisplay = false;
            continue;
        }
        flag = false;
        switch (time) {
            case 0:
                flag = true;
                break;
            case 1:
                if (wealth_result.wealth_contents[item].time <= 1) flag = true;
                break;
            case 2:
                if (wealth_result.wealth_contents[item].time <= 3 && wealth_result.wealth_contents[item].time > 1) flag = true;
                break;
            case 3:
                if (wealth_result.wealth_contents[item].time <= 6 && wealth_result.wealth_contents[item].time > 3) flag = true;
                break;
            case 4:
                if (wealth_result.wealth_contents[item].time <= 12 && wealth_result.wealth_contents[item].time > 6) flag = true;
                break;
            case 5:
                if (wealth_result.wealth_contents[item].time > 12) flag = true;
                break;
        }
        if (!flag) {
            wealth_result.wealth_contents[item].ifdisplay = false;
            continue;
        }
        flag = false;
        switch (profit) {
            case 0:
                flag = true;
                break;
            case 1:
                if (wealth_result.wealth_contents[item].profit <= 5) flag = true;
                break;
            case 2:
                if (wealth_result.wealth_contents[item].profit <= 10 && wealth_result.wealth_contents[item].profit > 5) flag = true;
                break;
            case 3:
                if (wealth_result.wealth_contents[item].profit <= 15 && wealth_result.wealth_contents[item].profit > 10) flag = true;
                break;
            case 4:
                if (wealth_result.wealth_contents[item].profit > 15) flag = true;
                break;
        }
        if (!flag) {
            wealth_result.wealth_contents[item].ifdisplay = false;
            continue;
        }
        wealth_result.wealth_contents[item].ifdisplay = true;
    }
};