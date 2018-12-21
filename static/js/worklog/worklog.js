!function(){
    laydate.skin('dahong');
    laydate({elem: '#demo'});
}();

function workLogload() {
    $('#table').bootstrapTable({
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
                workLogTitle: $("#workLogTitle").val(),
                workLogContent: $("#workLogContent").val(),
                projectId: $("#projectId").val()
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
            },
            {
                title: '更新时间',
                field: 'WorkLogUpdateTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.WorkLogUpdateTime.substr(0,10);
                }
            },
            {
                title: '操作',
                field: '',
                align: 'center',
                formatter: function (value, row) {
                    var detail = '<button button="#" mce_href="#" onclick="detailWorkLog(\'' + row.WorkLogId + '\')">详情</button> ';
                    var del = '<button button="#" mce_href="#" onclick="delWorkLog(\'' + row.WorkLogId + '\')">删除</button> ';
                    var edit = '<button button="#" mce_href="#" onclick="editWorkLog(\'' + row.WorkLogId + '\')">编辑</button> ';
                    return detail + del + edit;
                }
            }
        ]
    });
}

function getWorkLogTableData() {
    $('#table').bootstrapTable('refresh');
}

function addWorkLog() {
    var url = '/workLog/add'
    openlayer(url)
}

function editWorkLog(WorkLogId) {
    var url = '/workLog/update?WorkLogId='+WorkLogId
    openlayer(url)
}

function detailWorkLog(WorkLogId) {
    var url = '/workLog/detail?WorkLogId='+WorkLogId
    openlayer(url)
}

function delWorkLog(WorkLogId) {
    $.ajax({
        url: '/workLog/delete',
        type: 'post',
        data: {
            WorkLogId: WorkLogId
        },
        dataType: 'json',
        success: function (data) {
            if (data.status == 1) {
                alert("删除成功！")
                getWorkLogTableData();
            } else {
                alert("删除失败")
            }
        },
        error: function (err) {
        }
    });
}

function openlayer(url) {
    layer.open({
        type: 2,
        title: '工作日志信息',
        shadeClose: true,
        shade: 0.5,
        skin: 'layui-layer-rim',
        area: ['98%', '90%'],
        shadeClose: true,
        closeBtn: 2,
        content: url
    });
}





