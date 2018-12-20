function change() {
    var height01 = $(window).height();
    $(".top").css('height', height01 - 35+"px");
}

!function () {
    laydate.skin('dahong');
    laydate({ elem: '#CompanyInTime' });
    laydate({ elem: '#CompanyOutTime' });
}();

$(document).ready(function () {
    $("#form_company_add_save").click(saveOrUpdateCompany);
    if($("#CompanyOutTime").val() == "0001-01-01") {
        $("#CompanyOutTime").val("");
    }
});

function saveOrUpdateCompany() {
    if($("#CompanyId").val()) {
        updateCompany();
    } else {
        saveCompany();
    }
}

function saveCompany() {
    $.ajax({
        type: "POST",
        url: "/company/add",
        data: $("#form_company").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("保存成功！！！");
                parent.getCompanyTableData();
                TaskCancel();
            } else {
                alert("保存失败！！！")
            }
        }
    })
}

function updateCompany() {
    $.ajax({
        type: "POST",
        url: "/company/update",
        data: $("#form_company").serialize(),
        dataType: "json",
        success: function (result) {
            if (result.status == 1) {
                alert("更新成功！！！");
                parent.getCompanyTableData();
                TaskCancel();
            } else {
                alert("更新失败！！！")
            }
        }
    })
}

function TaskCancel() {
    var index = parent.layer.getFrameIndex(window.name)
    parent.layer.close(index);
}