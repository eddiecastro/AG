const Project=(function($){

    const Variables={
        IDSelectors:{
           ContactsTable:"#contacts-table",
           CounterTable:"#counter-table",
           DupesTable:"#dupes-table",
           BtnGetCounter:"#btn-get-counter",
           BtnShowDupes:"#btn-show-dupes",

        },
        ClassSelectors:{

        },
        DataSelectors:{

        },
        DataValueSelectors:{

        },
        ContactsTable:null,
        CounterTable:null,
        DupesTable:null,

    }

    const init = function(){
        LoadContactsTable();
        registerLoadWordCounter();
    }

    const LoadContactsTable = function(teamID){
        AjaxRequest("http://localhost:8000/people","GET",null,function(d){

            Variables.ContactsTable = $(Variables.IDSelectors.ContactsTable).DataTable({
                destroy: true,
                "stripeClasses": [],
                "lengthMenu": [5, 10, 20, 50],
                "pageLength": 5,
                data: d,
                columns: [
                    { data: 'id' },
                    { data: 'display_name' },
                    { data: 'email_address' },
                    { data: 'title' },
               ],
               "columnDefs":
                   [
                       {
                           "targets": 0,
                           "visible": false
                       },
                   ]
            });

        },function(d){
            console.log(d);
        });
    }

    const registerLoadWordCounter = function(){
        $(Variables.IDSelectors.BtnGetCounter).off('click.BtnGetCounter').on('click.BtnGetCounter',function(){
            LoadCounterTable();
        })
    }

    const LoadCounterTable = function(teamID){
        AjaxRequest("http://localhost:8000/people/wordscounter","GET",null,function(d){

            Variables.CounterTable = $(Variables.IDSelectors.CounterTable).DataTable({
                destroy: true,
                "stripeClasses": [],
                "lengthMenu": [5, 10, 20, 50],
                "pageLength": 5,
                data: d,
                columns: [
                    { data: 'word' },
                    { data: 'counter' },
               ],
            });

        },function(d){
            console.log(d);
        });
    }

    const AjaxRequest = function(url, method, data, funcSuccess, funError){
        $.ajax({
            url: url,
            type: method,
            dataType: 'json',
            contentType: "application/json; charset=utf-8",
            data: data,
            success: function (d) {
                funcSuccess(d);
            },
            error: function (d) {
                funError(d);
            }
        })  
    }

    return {
        init:init

    }

})(jQuery)