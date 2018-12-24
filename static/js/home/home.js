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

function newsLoad() {
    $.ajax({
        type: "get",
        url: "/news/list",
        dataType: "json",
        success: function (result) {
            var html = "";
            $.each(result, function(index, item){
                html += '<li class="ue-clear">'+
                            '<a href="'+item.NewsUrl+'" target="_blank" class="notice-title">'+item.NewsTitle+'</a>'+
                            '<div class="notice-time">'+item.NewsDate+'</div>'+
                        '</li>';
            });
            $("#newsUL").html(html);
        }
    })
}

function workLogLoad() {
    $('#workLogTable').bootstrapTable({
        method: "get",
        url: "/workLog/doList",
        striped: true,
        singleSelect: false,
        dataType: "json",
        pagination: true,
        sidePagination: "server",
        pageSize: 10,
        pageNumber: 1,
        search: false,
        queryParams: function(params) {
            var temp = {
                pageSize: params.limit,
                pageNo: (params.offset / params.limit) + 1,
                projectId: -1
            }
            return temp
        },
        columns: [
            {
                title: '编号',
                field: 'WorkLogId',
                align: 'center',
                valign: 'middle',
                hidden: true
            },
            {
                title: '工作日志名称',
                field: 'WorkLogTitle',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '项目名称',
                field: 'Project.ProjectName',
                align: 'center'
            },
            {
                title: '记录时间',
                field: 'WorkLogCreateTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.WorkLogCreateTime.substr(0,10);
                }
            }
        ]
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
                pageNo: (params.offset / params.limit) + 1
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
    newsLoad();
    workLogLoad();
    projectLoad();
}

$(document).ready(function() {
    tableLoad();
});