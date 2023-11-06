<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
    export let listCurr = [];
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "COMPANY"
	let title_idcompany = ""
	let title_minbet = ""
    let sData = "";
    let sDataListBet = "";
    let sDataConfPoint = "";
    let myModal_newentry = "";
    let flag_code = false;
    let flag_btnsave = true;
    let name_field = "";
    let idcurr_field = "";
    let url_field = "";
    let owner_field = "";
    let phone_field = "";
    let email_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";
    let idrecord = "";
    let searchHome = "";
    let filterHome = [];

    //COMPANY INVOICE
    let invoice_company = [];
    let invoice_id = "";
    let invoice_username = "";
    let invoice_roundbet = 0;
    let invoice_totalbet = 0;
    let invoice_totalwin = 0;
    let invoice_totalbonus = 0;
    let invoice_card_codepoin = "";
    let invoice_card_pattern = "";
    let invoice_card_pattern_sort = "";
    let invoice_card_result = "";
    let invoice_card_win = "";

    //COMPANY LISTBET
    let listbet_master = [];
    let listbet_company = [];
    let listadmin_company = [];
    let idrecord_listbet_field = 0;
    let idcompany_listbet_field = "";
    let minbet_listbet_field = 0;

     //COMPANY CONF POINT
    let listconfpoint_company = [];
    let idrecord_confpoint_field = 0;
    let idbet_confpoint_field = 0;
    let idcompany_confpoint_field = "";
    let point_confpoint_field = 0;
    let nmpoint_confpoint_field = "";
    let create_confpoint_field = "";
    let update_confpoint_field = "";


    let css_loader = "display: none;";
    let msgloader = "";
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,name,idcurr,url,owner,email,phone,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            if(status == "N"){
                flag_btnsave = false;
            }else{
                flag_btnsave = true;
            }
            flag_code = true;
            idrecord = id
            name_field = name;
            idcurr_field = idcurr;
            url_field = url;
            owner_field = owner;
            phone_field = phone;
            email_field = email;
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const EditConfPoint = (id,idbet,nmpoin,poin,create,update) => {
        sDataListBet = "Edit"
        idrecord_confpoint_field = id
        idbet_confpoint_field = idbet
        nmpoint_confpoint_field = nmpoin
        point_confpoint_field = poin
        create_confpoint_field = create;
        update_confpoint_field = update;
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalformconfpoint"));
        myModal_newentry.show();
        
    };
    const call_listadmin = (e) => {
        title_idcompany = e
        call_listadmin_bycompany(e)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistadmin"));
        myModal_newentry.show();
        
    };
    const call_listconfpoint = (e,f) => {
        sDataConfPoint = "New"
        title_idcompany = idcompany_listbet_field
        title_minbet = f
        idbet_confpoint_field = e
        idcompany_confpoint_field = idcompany_listbet_field
        point_confpoint_field = 0;
        call_listconfpoint_bycompany(e)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistconfpoint"));
        myModal_newentry.show();
        
    };
    const call_invoice = (e) => {
        title_idcompany = e
        call_invoice_bycompany(e)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalinvoice"));
        myModal_newentry.show();
        
    };
    const call_invoicedetail = (code,username,roundbet,totalbet,totalwin,totalbonus,card_codepoin,card_pattern,card_result,card_win) => {
        invoice_id = code;
        invoice_username = username;
        invoice_roundbet = roundbet;
        invoice_totalbet = totalbet;
        invoice_totalwin = totalwin
        invoice_totalbonus = totalbonus
        invoice_card_codepoin = card_codepoin;
        invoice_card_pattern = card_pattern;
        invoice_card_result = card_result;
        invoice_card_win = card_win;

        let shuffleArray = []
        let temp_data = invoice_card_result.split('-')
        let temp_data_total = temp_data.length
        invoice_card_pattern_sort = ""
        for(let i=0;i<temp_data_total;i++) {
            shuffleArray.push(card_result_data[temp_data[i]])
        }
        console.log(invoice_card_result)
        console.log(shuffleArray)
        function compareByval_display(a, b) {
            return a.val_display - b.val_display;
          }
        let pattern_array_sort = shuffleArray.sort(compareByval_display);
        for(let i=0;i<pattern_array_sort.length;i++) {
        
            if(i == pattern_array_sort.length-1) {
                let temp_data = card_result_data.findIndex(card => card.id ==pattern_array_sort[i].id)
                invoice_card_pattern_sort += temp_data
            }else{
                let temp_data = card_result_data.findIndex(card => card.id ==pattern_array_sort[i].id)
                invoice_card_pattern_sort += temp_data+"-"
            }
        }
        console.log(invoice_card_pattern_sort)

        myModal_newentry = new bootstrap.Modal(document.getElementById("modalinvoicedetail"));
        myModal_newentry.show();
       
    };
    const call_listbet = (e) => {
        sDataListBet = "New"
        title_idcompany = e
        idrecord_listbet_field = 0;
        idcompany_listbet_field = e;
        minbet_listbet_field = 0;
        call_listbet_master(e)
        call_listbet_bycompany(e)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistbet"));
        myModal_newentry.show();
        
    };
    const call_formlistbet = () => {
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalcrudlistbetcompany"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idrecord.length == 4){
                flag = false
                msg += "The ID is maxlengt 4\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companysave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    company_id: idrecord.toUpperCase(),
                    company_name: name_field,
                    company_idcurr:idcurr_field,
                    company_nmowner: owner_field,
                    company_phoneowner: phone_field,
                    company_emailowner: email_field,
                    company_url: url_field,
                    company_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function call_invoice_bycompany(idcompany) {
        invoice_company = [];
        const res = await fetch("/api/companyinvoice", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company_id: idcompany.toLowerCase(),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                let winlose = 0;
                let winlose_css = "";
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    winlose = (parseInt(record[i]["companyinvoice_totalbet"]) * parseInt(record[i]["companyinvoice_roundbet"])) - (parseInt(record[i]["companyinvoice_totalwin"])+parseInt(record[i]["companyinvoice_totalbonus"]))
                    
                    if(winlose > 0){
                        winlose_css = "color:blue;font-weight:bold;"
                    }else{
                        winlose_css = "color:red;font-weight:bold;"
                    }
                    let temp_string = record[i]["companyinvoice_create"]
                    let temp_result = temp_string.split(",");
                    invoice_company = [
                        ...invoice_company,
                        {
                        companyinvoice_no: no,
                        companyinvoice_id: record[i]["companyinvoice_id"],
                        companyinvoice_username: record[i]["companyinvoice_username"],
                        companyinvoice_date: temp_result[1],
                        companyinvoice_roundbet: record[i]["companyinvoice_roundbet"],
                        companyinvoice_totalbet: record[i]["companyinvoice_totalbet"],
                        companyinvoice_totalwin: record[i]["companyinvoice_totalwin"],
                        companyinvoice_totalbonus: record[i]["companyinvoice_totalbonus"],
                        companyinvoice_winlose: winlose,
                        companyinvoice_winlose_css: winlose_css,
                        companyinvoice_card_codepoin: record[i]["companyinvoice_card_codepoin"],
                        companyinvoice_card_pattern: record[i]["companyinvoice_card_pattern"],
                        companyinvoice_card_result: record[i]["companyinvoice_card_result"],
                        companyinvoice_card_win: record[i]["companyinvoice_card_win"],
                        companyinvoice_status: record[i]["companyinvoice_status"],
                        companyinvoice_status_css: record[i]["companyinvoice_status_css"],
                        companyinvoice_create: record[i]["companyinvoice_create"],
                        companyinvoice_update: record[i]["companyinvoice_update"],
                        },
                    ];
                }
            }
        }
    }
    async function call_listconfpoint_bycompany(idbet) {
        listconfpoint_company = [];
        const res = await fetch("/api/companyconfpoint", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company_idbet: parseInt(idbet),
                company_id: idcompany_listbet_field,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listconfpoint_company = [
                        ...listconfpoint_company,
                        {
                        companyconf_no: no,
                        companyconf_id: record[i]["companyconf_id"],
                        companyconf_idbet: record[i]["companyconf_idbet"],
                        companyconf_idpoin: record[i]["companyconf_idpoin"],
                        companyconf_nmpoin: record[i]["companyconf_nmpoin"],
                        companyconf_poin: record[i]["companyconf_poin"],
                        companyconf_create: record[i]["companyconf_create"],
                        companyconf_update: record[i]["companyconf_update"],
                        },
                    ];
                }
            }
        }
    }
    async function call_listbet_master() {
        listbet_master = [];
        const res = await fetch("/api/listbet", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
               
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listbet_master = [
                        ...listbet_master,
                        {
                        lisbet_id: record[i]["lisbet_id"],
                        lisbet_minbet: record[i]["lisbet_minbet"],
                        },
                    ];
                }
            }
        }
    }
    async function call_listbet_bycompany(idcompany) {
        listbet_company = [];
        const res = await fetch("/api/companylistbet", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company_id: idcompany,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listbet_company = [
                        ...listbet_company,
                        {
                        companylistbet_no: no,
                        companylistbet_id: record[i]["companylistbet_id"],
                        companylistbet_minbet: record[i]["companylistbet_minbet"],
                        companylistbet_create: record[i]["companylistbet_create"],
                        companylistbet_update: record[i]["companylistbet_update"],
                        },
                    ];
                }
            }
        }
    }
    async function call_listadmin_bycompany(idcompany) {
        listadmin_company = [];
        const res = await fetch("/api/companyadminbycompany", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                company_id: idcompany,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listadmin_company = [
                        ...listadmin_company,
                        {
                        companyadmin_no: no,
                        companyadmin_id: record[i]["companyadmin_id"],
                        companyadmin_name: record[i]["companyadmin_name"],
                        companyadmin_tipe: record[i]["companyadmin_tipe"],
                        companyadmin_status: record[i]["companyadmin_status"],
                        companyadmin_status_css: record[i]["companyadmin_status_css"],
                        companyadmin_create: record[i]["companyadmin_create"],
                        companyadmin_update: record[i]["companyadmin_update"],
                        },
                    ];
                }
            }
        }
    }
    async function handleSave_listbet() {
        let flag = true
        let msg = ""
        let min_bet = value_listbetmaster(minbet_listbet_field)
        if(sDataListBet == "New"){
            if(idcompany_listbet_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(minbet_listbet_field == ""){
                flag = false
                msg += "The Minbet is required\n"
            }
            if(parseInt(min_bet) <= 0){
                flag = false
                msg += "The Minbet cannot 0\n"
            }
        }else{
            if(idrecord_listbet_field == "" || idrecord_listbet_field == 0){
                flag = false
                msg += "The ID is required\n"
            }
            if(idcompany_listbet_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(minbet_listbet_field == ""){
                flag = false
                msg += "The Minbet is required\n"
            }
            if(parseInt(min_bet) <= 0){
                flag = false
                msg += "The Minbet cannot 0\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companylistbetsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataListBet,
                    page:"CURR-SAVE",
                    companylistbet_id: parseInt(idrecord_listbet_field),
                    companylistbet_idcompany: idcompany_listbet_field,
                    companylistbet_minbet:parseInt(min_bet),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField_listbet()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleSave_confpoint_generate() {
        let flag = true
        let msg = ""
        if(sDataConfPoint == "New"){
            if(idcompany_confpoint_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(idbet_confpoint_field == ""){
                flag = false
                msg += "The IDBET is required\n"
            }
        }else{
            if(idrecord_confpoint_field == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idcompany_confpoint_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(idbet_confpoint_field == ""){
                flag = false
                msg += "The IDBET is required\n"
            }
            if(point_confpoint_field == ""){
                flag = false
                msg += "The Point is required\n"
            }
            if(point_confpoint_field <= 0){
                flag = false
                msg += "The Point cannot 0\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/companyconfpointsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataListBet,
                    page:"CURR-SAVE",
                    companyconfpoint_id: parseInt(idrecord_confpoint_field),
                    companyconfpoint_idbet: parseInt(idbet_confpoint_field),
                    Companyconfpoint_idcompany: idcompany_confpoint_field,
                    Companyconfpoint_point:parseInt(point_confpoint_field),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField_listbet()
                }
                msgloader = json.message;
                call_listconfpoint_bycompany(idbet_confpoint_field)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    const card_result_data = [
      {id:"2_diamond",val:"2",val_display:2,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_2.png"},
      {id:"3_diamond",val:"3",val_display:3,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_3.png"},
      {id:"4_diamond",val:"4",val_display:4,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_4.png"},
      {id:"5_diamond",val:"5",val_display:5,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_5.png"},
      {id:"6_diamond",val:"6",val_display:6,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_6.png"},
      {id:"7_diamond",val:"7",val_display:7,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_7.png"},
      {id:"8_diamond",val:"8",val_display:8,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_8.png"},
      {id:"9_diamond",val:"9",val_display:9,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_9.png"},
      {id:"10_diamond",val:"10",val_display:10,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_10.png"},
      {id:"j_diamond",val:"J",val_display:11,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_J.png"},
      {id:"q_diamond",val:"Q",val_display:12,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_Q.png"},
      {id:"k_diamond",val:"K",val_display:13,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_K.png"},
      {id:"as_diamond",val:"AS",val_display:14,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_AS.png"},
      {id:"2_heart",val:"2",val_display:2,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_2.png"},
      {id:"3_heart",val:"3",val_display:3,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_3.png"},
      {id:"4_heart",val:"4",val_display:4,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_4.png"},
      {id:"5_heart",val:"5",val_display:5,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_5.png"},
      {id:"6_heart",val:"6",val_display:6,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_6.png"},
      {id:"7_heart",val:"7",val_display:7,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_7.png"},
      {id:"8_heart",val:"8",val_display:8,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_8.png"},
      {id:"9_heart",val:"9",val_display:9,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_9.png"},
      {id:"10_heart",val:"10",val_display:10,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_10.png"},
      {id:"j_heart",val:"J",val_display:11,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_J.png"},
      {id:"q_heart",val:"Q",val_display:12,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_Q.png"},
      {id:"k_heart",val:"K",val_display:13,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_K.png"},
      {id:"as_heart",val:"AS",val_display:14,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_AS.png"},
      {id:"2_club",val:"2",val_display:2,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_2.png"},
      {id:"3_club",val:"3",val_display:3,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_3.png"},
      {id:"4_club",val:"4",val_display:4,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_4.png"},
      {id:"5_club",val:"5",val_display:5,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_5.png"},
      {id:"6_club",val:"6",val_display:6,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_6.png"},
      {id:"7_club",val:"7",val_display:7,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_7.png"},
      {id:"8_club",val:"8",val_display:8,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_8.png"},
      {id:"9_club",val:"9",val_display:9,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_9.png"},
      {id:"10_club",val:"10",val_display:10,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_10.png"},
      {id:"j_club",val:"J",val_display:11,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_J.png"},
      {id:"q_club",val:"Q",val_display:12,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_Q.png"},
      {id:"k_club",val:"K",val_display:13,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_K.png"},
      {id:"as_club",val:"AS",val_display:14,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_AS.png"},
      {id:"2_spade",val:"2",val_display:2,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_2.png"},
      {id:"3_spade",val:"3",val_display:3,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_3.png"},
      {id:"4_spade",val:"4",val_display:4,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_4.png"},
      {id:"5_spade",val:"5",val_display:5,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_5.png"},
      {id:"6_spade",val:"6",val_display:6,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_6.png"},
      {id:"7_spade",val:"7",val_display:7,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_7.png"},
      {id:"8_spade",val:"8",val_display:8,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_8.png"},
      {id:"9_spade",val:"9",val_display:9,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_9.png"},
      {id:"10_spade",val:"10",val_display:10,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_10.png"},
      {id:"j_spade",val:"J",val_display:11,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_J.png"},
      {id:"q_spade",val:"Q",val_display:12,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_Q.png"},
      {id:"k_spade",val:"K",val_display:13,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_K.png"},
      {id:"as_spade",val:"AS",val_display:14,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_AS.png"},
      {id:"jk_black",val:"JK",val_display:1,code_card:"JK",img:"./CARD/WHITE/CARD_JOKER_BLACK.png"},
      {id:"jk_red",val:"JK",val_display:1,code_card:"JK",img:"./CARD/WHITE/CARD_JOKER_RED.png"},
    ]
    const list_point = [
      {id:"1",code:"RF",name:"Royal Flush",poin:500},
      {id:"2",code:"5K",name:"5 Of A Kind",poin:200},
      {id:"3",code:"SF",name:"Straight Flush",poin:120},
      {id:"4",code:"4K",name:"4 Of A Kind",poin:50},
      {id:"5",code:"FH",name:"Full House",poin:7},
      {id:"6",code:"FL",name:"Flush",poin:5},
      {id:"7",code:"ST",name:"Straight",poin:3},
      {id:"8",code:"3K",name:"3 Of A Kind",poin:2},
      {id:"9",code:"2P",name:"2 Pair (10 PAIR)",poin:1},
      {id:"10",code:"AP",name:"Ace Pair",poin:1},
    ]
    function card_img(e,y){ 
      if(e != "" || e.length > 0){
        let data = e.split("-");
        let total_data = e.split("-").length;
        let img_data = "";
        for(let i=0;i<total_data;i++){
        //   const searchIndex = card_result_data.findIndex((car) => car.id==data[i]);
          
          img_data +="<img width='"+y+"px' src='"+card_result_data[data[i]].img+"' /> "
        }
        return img_data
      }else{
        return ""
      }
    }
    function card_img2(e,y){ 
       
      if(e != "" || e.length > 0){
        let data = e.split(",");
        let total_data = e.split(",").length;
        let img_data = "";
        for(let i=0;i<total_data;i++){
          const searchIndex = card_result_data.findIndex((car) => car.id==data[i]);
          
          img_data +="<img width='"+y+"px' src='"+card_result_data[searchIndex].img+"' /> "
        }
        return img_data
      }else{
        return ""
      }
    }
    function card_codepoin(e){
        if(e != ""){
            const searchIndex = list_point.findIndex((car) => car.code==e);
            return list_point[searchIndex]["name"]
        }else{
            return ""
        } 
        
    }
    function clearField(){
        flag_code = false
        idrecord = "";
        name_field = "";
        idcurr_field = "";
        url_field = "";
        owner_field = "";
        phone_field = "";
        email_field = "";
        status_field = "";
        create_field = "";
        update_field = "";
    }
    function clearField_listbet(){
        flag_code = false
        idrecord_listbet_field = 0;
        idcompany_listbet_field = "";
        minbet_listbet_field = 0;
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
        }  
    };
    function uperCase(element) {
        function onInput(event) {
            element.value = element.value.toUpperCase();
            element.value = element.value.replace(/[^A-Z0-9]/gi, '');
        }
        element.addEventListener("input", onInput);
        return {
            destroy() {
                element.removeEventListener("input", onInput);
            },
        };
    }
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
    }
    function value_listbetmaster(e){
        let value = 0
        for(let i=0;i<listbet_master.length;i++){
            if(e == listbet_master[i].lisbet_id){
                value = parseInt(listbet_master[i].lisbet_minbet)
            }
        }
        return value;
    }
    const handleKeyboard_int = () => {
        let numbera;
		for (let i = 0; i < point_confpoint_field.length; i++) {
			numbera = parseInt(point_confpoint_field[i]);
			if (isNaN(numbera)) {
				point_confpoint_field = "";
			}
		}
	}
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Code"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan=4>&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CODE</th>
                                <th NOWRAP width="2%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CURR</th>
                                <th NOWRAP width="7%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">START</th>
                                <th NOWRAP width="7%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">END</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">COMPANY</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">OWNER</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">EMAIL</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                //e,id,name,idcurr,url,owner,email,phone,status,create,update
                                                NewData("Edit",rec.home_id, 
                                                rec.home_name, rec.home_idcurr,rec.home_url,
                                                rec.home_nmowner, rec.home_emailowner, 
                                                rec.home_phoneowner,rec.home_status,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                call_invoice(rec.home_id)
                                            }} class="bi bi-file-earmark"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                call_listbet(rec.home_id)
                                            }} class="bi bi-file-earmark"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                call_listadmin(rec.home_id)
                                            }} class="bi bi-people"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_idcurr}</td>
                                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_start}</td>
                                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_end}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <a href="{rec.home_url}" target="_blank">{rec.home_name}</a>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmowner}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_phoneowner}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_emailowner}</td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">CODE</label>
                    <input bind:value={idrecord}
                        use:uperCase  
                        disabled={flag_code}
                        class="required form-control"
                        maxlength="9"
                        type="text"
                        placeholder="CODE"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Default Currency</label>
                    <select
                        bind:value="{idcurr_field}" 
                        name="currency" id="currency" 
                        class="required form-control">
                        {#each listCurr as rec}
                        <option value="{rec.curr_id}">{rec.curr_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Company</label>
                    <Input bind:value={name_field}
                        class="required"
                        type="text"
                        placeholder="Company"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">URL</label>
                    <Input bind:value={url_field}
                        class="required"
                        type="text"
                        placeholder="URL"/>
                </div>
            </div>
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Owner</label>
                    <Input bind:value={owner_field}
                        class="required"
                        type="text"
                        placeholder="Owner"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone</label>
                    <Input bind:value={phone_field}
                        class="required"
                        type="text"
                        placeholder="Phone"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Email</label>
                    <Input bind:value={email_field}
                        class=""
                        type="text"
                        placeholder="Email"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                    <div class="mb-3">
                        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                            Create : {create_field}<br />
                            Update : {update_field}
                        </div>
                    </div>
                {/if}
            </div>
        </div>
        
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modallistbet"
	modal_size="modal-dialog-centered"
	modal_title="ListBet - {title_idcompany}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="*" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">MINBET</th>
                </tr>
            </thead>
            <tbody>
                {#each listbet_company as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                            <i on:click={() => {
                                    //e,id,name,idcurr,url,owner,email,phone,status,create,update
                                    call_listconfpoint(rec.companylistbet_id,rec.companylistbet_minbet);
                                }} class="bi bi-pencil"></i>
                        </td>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.companylistbet_no}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};">{new Intl.NumberFormat().format(rec.companylistbet_minbet)}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
                call_formlistbet();
            }} 
            button_title="New"
            button_css="btn-info"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalcrudlistbetcompany"
	modal_size="modal-dialog-centered"
	modal_title="{title_page+" "+title_idcompany+" / "+sDataListBet}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Min Bet</label>
            <select
                class="form-control required"
                style="text-align: right;"
                bind:value={minbet_listbet_field}>
                {#each listbet_master as rec}
                 <option value="{rec.lisbet_id}">{new Intl.NumberFormat().format(rec.lisbet_minbet)}</option>
                {/each}
            </select>
        </div>
        {#if sDataListBet != "New"}
            <div class="mb-3">
                <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                    Create : {create_field}<br />
                    Update : {update_field}
                </div>
            </div>
        {/if}
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave_listbet();
            }} 
            
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modallistconfpoint"
	modal_size="modal-dialog-centered"
	modal_title="Configure Point - {title_idcompany} - {title_minbet}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">POINT</th>
                </tr>
            </thead>
            <tbody>
                {#each listconfpoint_company as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                            <i on:click={() => {
                                    //e,id,name,idcurr,url,owner,email,phone,status,create,update
                                    EditConfPoint(rec.companyconf_id,rec.companyconf_idbet,rec.companyconf_nmpoin,rec.companyconf_poin,
                                    rec.companyconf_create,rec.companyconf_update);
                                }} class="bi bi-pencil"></i>
                        </td>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.companyconf_no}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyconf_nmpoin}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};">{new Intl.NumberFormat().format(rec.companyconf_poin)}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
                handleSave_confpoint_generate();
            }} 
            button_title="Generate"
            button_css="btn-info"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalformconfpoint"
	modal_size="modal-dialog-centered "
	modal_title="Update Configure Point"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <input bind:value={nmpoint_confpoint_field}
                disabled
                class="required form-control"
                type="text"
                placeholder="Name"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Point</label>
            <Input
                on:keyup={handleKeyboard_int} 
                bind:value={point_confpoint_field}
                style="text-align: right;"
                class="required"
                type="text"
                placeholder="Point"/>
            <div style="text-align: right; color:blue;font-size:11px;">
                {new Intl.NumberFormat().format(point_confpoint_field)}
            </div>
        </div>
        {#if sData != "New"}
            <div class="mb-3">
                <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                    Create : {create_confpoint_field}<br />
                    Update : {update_confpoint_field}
                </div>
            </div>
        {/if}
      
        
        
        
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave_confpoint_generate();
            }} 
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modallistadmin"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="ListAdmin - {title_idcompany}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TIPE</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">USERNAME</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NAME</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CREATE</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UPDATE</th>
                </tr>
            </thead>
            <tbody>
                {#each listadmin_company as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_no}</td>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                            <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.companyadmin_status_css}">
                                {status(rec.companyadmin_status)}
                            </span>
                        </td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_tipe}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_id}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_name}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_create}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyadmin_update}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
	</slot:template>
</Modal>

<Modal
	modal_id="modalinvoice"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="Invoice - {title_idcompany}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">INVOICE</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">USERNAME</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DATE</th>
                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PATTERN</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ROUND</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BET</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TOTALBET</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TOTALWIN</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TOTALBONUS</th>
                    <th NOWRAP width="50%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">WINLOSE</th>
                </tr>
            </thead>
            <tbody>
                {#each invoice_company as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.companyinvoice_no}</td>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                            <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.companyinvoice_status_css}">{rec.companyinvoice_status}</span>
                        </td>
                        <td on:click={() => {
                            //code,username,totalbet,totalwin,totalbonus,card_codepoin,card_pattern,card_result,card_win
                            call_invoicedetail(rec.companyinvoice_id,rec.companyinvoice_username,rec.companyinvoice_roundbet,
                            rec.companyinvoice_totalbet,rec.companyinvoice_totalwin,rec.companyinvoice_totalbonus,
                            rec.companyinvoice_card_codepoin,rec.companyinvoice_card_pattern,rec.companyinvoice_card_result,rec.companyinvoice_card_win);
                        }} NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};top;cursor:pointer;text-decoration:underline;">
                            {rec.companyinvoice_id}
                        </td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyinvoice_username}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.companyinvoice_date}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                            {rec.companyinvoice_card_pattern}<br />
                            {rec.companyinvoice_card_result}  {card_codepoin(rec.companyinvoice_card_codepoin)}
                        </td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{rec.companyinvoice_roundbet}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(rec.companyinvoice_totalbet)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;font-weight:bold;">{new Intl.NumberFormat().format(rec.companyinvoice_roundbet*rec.companyinvoice_totalbet)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;font-weight:bold;">{new Intl.NumberFormat().format(rec.companyinvoice_totalwin)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;font-weight:bold;">{new Intl.NumberFormat().format(rec.companyinvoice_totalbonus)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.companyinvoice_winlose_css}">{new Intl.NumberFormat().format(rec.companyinvoice_winlose)}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        
	</slot:template>
</Modal>

<Modal
	modal_id="modalinvoicedetail"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="Invoice - {title_idcompany}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table style="width: 100%;">
            <tr>
                <td width="50%" style="font-size: 12px;">INVOICE</td>
                <td width="1%" style="font-size: 12px;">:</td>
                <td width="*" style="font-size: 12px;">{invoice_id}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">USERNAME</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: left;">{invoice_username}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">TOTAL BET</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;color:blue;font-weight: bold;">{invoice_roundbet}*{new Intl.NumberFormat().format(invoice_totalbet)} = {new Intl.NumberFormat().format(invoice_roundbet*invoice_totalbet)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">TOTAL WIN</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;color:red;font-weight: bold;">{new Intl.NumberFormat().format(invoice_totalwin)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">TOTAL BONUS</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;color:red;font-weight: bold;">{new Intl.NumberFormat().format(invoice_totalbonus)}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">TOTAL WINLOSE</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: right;color:red;font-weight: bold;">
                    {new Intl.NumberFormat().format((invoice_roundbet*invoice_totalbet)-(invoice_totalbonus+invoice_totalwin))}
                </td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">CARD PATTERN</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: left;">{invoice_card_pattern}</td>
            </tr>
            <tr>
                <td style="font-size: 12px;text-align: left;">CARD POIN</td>
                <td style="font-size: 12px;text-align: left;">:</td>
                <td style="font-size: 12px;text-align: left;">{card_codepoin(invoice_card_codepoin)}</td>
            </tr>
        </table>
        CARD RESULT<br />
        {@html card_img(invoice_card_result,85)}<br /><br />
        CARD RESULT - SORT<br />
        {@html card_img(invoice_card_pattern_sort,85)}<br /><br />
        CARD WIN<br />
        {@html card_img2(invoice_card_win,85)}<br /><br />
	</slot:template>
	<slot:template slot="footer">
        
	</slot:template>
</Modal>