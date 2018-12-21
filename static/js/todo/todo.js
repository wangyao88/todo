!function(){
    laydate.skin('dahong');
    laydate({elem: '#TodoStartTime'});
    laydate({elem: '#TodoEndTime'});
}();

function todoload() {
    $('#table').bootstrapTable({
        method: "get",
        url: "/todo/doList",
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
                todoTitle: $("#todoTitle").val(),
                todoContent: $("#todoContent").val(),
                todoStartTime: $("#TodoStartTime").val(),
                todoEndTime: $("#TodoEndTime").val()
            }
            return temp
        },
        columns: [
            {
                title: '编号',
                field: 'TodoId',
                align: 'center',
                valign: 'middle',
                hidden: true
            },
            {
                title: '待办事项名称',
                field: 'TodoTitle',
                align: 'center',
                valign: 'middle'
            },
            {
                title: '待办事项内容',
                field: 'TodoContent',
                align: 'center'
            },
            {
                title: '创建时间',
                field: 'TodoCreateTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.TodoCreateTime.substr(0,10);
                }
            },
            {
                title: '预计完成时间',
                field: 'TodoExcpteEndTime',
                align: 'center',
                formatter: function (value, row) {
                    return row.TodoExcpteEndTime.substr(0,10);
                }
            },
            {
                title: '实际完成时间',
                field: 'TodoRealyEndTime',
                align: 'center',
                formatter: function (value, row) {
                    if(row.CompanyOutTime == "0001-01-01T00:00:00Z") {
                        return "";
                    }
                    var subTime = row.TodoRealyEndTime.substr(0,10);
                    if(subTime == "0001-01-01") {
                        return "";
                    }
                    return subTime
                }
            },
            {
                title: '操作',
                field: '',
                align: 'center',
                formatter: function (value, row) {
                    var del = '<button button="#" mce_href="#" onclick="delTodo(\'' + row.TodoId + '\')">删除</button> ';
                    var edit = '<button button="#" mce_href="#" onclick="editTodo(\'' + row.TodoId + '\')">编辑</button> ';
                    return del + edit;
                }
            }
        ]
    });
}

function getTodoTableData() {
    $('#table').bootstrapTable('refresh');
}

function addTodo() {
    var url = '/todo/add'
    openlayer(url)
}

function editTodo(TodoId) {
    var url = '/todo/update?TodoId='+TodoId
    openlayer(url)
}

function detailTodo(TodoId) {
    var url = '/todo/detail?TodoId='+TodoId
    openlayer(url)
}

function delTodo(TodoId) {
    $.ajax({
        url: '/todo/delete',
        type: 'post',
        data: {
            TodoId: TodoId
        },
        dataType: 'json',
        success: function (data) {
            if (data.status == 1) {
                alert("删除成功！")
                getTodoTableData();
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





