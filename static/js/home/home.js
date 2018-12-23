$(".title-list ul").on("click","li",function(){
    var aIndex = $(this).index();
    $(this).addClass("current").siblings().removeClass("current");
    $(".matter-content").removeClass("current").eq(aIndex).addClass("current");
});
$(".duty").find("tbody").find("tr:even").css("backgroundColor","#eff6fa");

function openlayer(id){
    layer.open({
        type: 2,
        title: '万年历',
        shadeClose: true,
        shade: 0.5,
        skin: 'layui-layer-rim',
        closeBtn:2,
        area: ['80%', '85%'],
        shadeClose: true,
        closeBtn: 2,
        content: '/calendar/list'
    });
}

function projectLoad() {
    $('#projectTable').bootstrapTable({
        method: "get",
        url: "/project/doList",
        striped: true,
        singleSelect: false,
        dataType: "json",
        pagination: true,
        sidePagination: "server",
        pageSize: 3,
        pageNumber: 1,
        pageList: [3, 10, 50, 100],
        search: false,
        queryParams: function(params) {
            var temp = {
                pageSize: params.limit,
                pageNo: (params.offset / params.limit) + 1,
                projectName: $("#projectName").val()
            }
            return temp
        },
        columns: [
            {
                title: '编号',
                field: 'ProjectId',
                align: 'center',
                valign: 'middle',
                hidden: true
            },
            {
                title: '项目名称',
                field: 'ProjectName',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '开始时间',
                field: 'ProjectStartTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.ProjectStartTime.substr(0,10);
                }
            },
            {
                title: '结束时间',
                field: 'ProjectEndTime',
                align: 'center',
                formatter: function (value, row) {
                    if (row.ProjectEndTime == "0001-01-01T00:00:00Z") {
                        return "至今";
                    }
                    return row.ProjectEndTime.substr(0, 10);
                }
            }
        ]
    });
}

function tableLoad() {
    projectLoad();
}

$(document).ready(function() {
    tableLoad();
});