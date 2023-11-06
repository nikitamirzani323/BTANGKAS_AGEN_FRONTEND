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
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "LIST PATTERN"
    let sData = "";
    let sDataListPatternDetail = "";
    let myModal_newentry = "";
    let flag_code = false;
    let flag_btnsave = true;
    let code_field = "";
    let name_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";

    let listpoint = []
    let listpatterndetail = []
    let idlistpattern_listpatterndetail_field = ""
    let status_listpatterndetail_field = "N"
    let flag_idpoin_field = false
    let idpoin_listpatterndetail_field = "0"
    let create_listpatterndetail_field = "";
    let update_listpatterndetail_field = "";
    let minbet_listpatterndetail = 500;
    let sumbet_listpatterndetail = 0;
    let sumwin_1_listpatterndetail = 0;
    let sumwin_2_listpatterndetail = 0;
    let sumwin_3_listpatterndetail = 0;
    let sumwin_4_listpatterndetail = 0;
    let sumwin_1_listpatterndetail_css = "color:red;";
    let sumwin_2_listpatterndetail_css = "color:red;";
    let sumwin_3_listpatterndetail_css = "color:red;";
    let sumwin_4_listpatterndetail_css = "color:red;";
    let totallose_listpatterndetail = 0;
    let totalwin_listpatterndetail = 0;
    let totallose_listpatterndetail_css = "color:red;";
    let totalwin_listpatterndetail_css = "color:red;";

    let idrecord = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_code
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_nmpoin
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id,name,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            flag_code = true;
            idrecord = id
            name_field = name;
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const Detaillistpattern = (e) => {
        sDataListPatternDetail = "New"
        idlistpattern_listpatterndetail_field = e
        call_listpatterndetail(e)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modallistpattern"));
        myModal_newentry.show();
        
    };
    const Deletelistpatterndetail = (e,y) => {
        handleListDetailPatternDeleteSingle(e,y)
    }
    const call_formlistpatterndetail = () => {
        clearField_listpatterndetail()
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalcrudlistdetailcompany"));
        myModal_newentry.show();
        
    };
    const handleSelectStatus = (event) => {
      if(event.target.value == "Y"){
        flag_idpoin_field = true
        call_listpoint()
      }else{
        idpoin_listpatterndetail_field = 0;
        flag_idpoin_field = false
      }
     
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    async function call_listpoint() {
        listpoint = [];
        const res = await fetch("/api/listpointshare", {
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
                    listpoint = [
                        ...listpoint,
                        {
                            lispoint_id: record[i]["lispoint_id"],
                            lispoint_name: record[i]["lispoint_code"]+"-"+record[i]["lispoint_name"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function call_listpatterndetail(e) {
        listpatterndetail = [];
        const res = await fetch("/api/listpatterndetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                listpatterndetail_idlistpattern:e
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            totallose_listpatterndetail = json.totallose;
            totalwin_listpatterndetail = json.totalwin;
            if(totallose_listpatterndetail > 0){
                totallose_listpatterndetail_css = "color:blue;";
            }else{
                totallose_listpatterndetail_css = "color:red;";
            }
            if(totalwin_listpatterndetail > 0){
                totalwin_listpatterndetail_css = "color:blue;";
            }else{
                totalwin_listpatterndetail_css = "color:red;";
            }
            if (record != null) {
                let no = 0;
                let nm_poin = "";
                let css_poin = "color:red;";
                sumbet_listpatterndetail = 0;
                sumwin_1_listpatterndetail = 0;
                sumwin_2_listpatterndetail = 0;
                sumwin_3_listpatterndetail = 0;
                sumwin_4_listpatterndetail = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if(record[i]["listpatterndetail_idpoin"] > 0){
                        nm_poin = record[i]["listpatterndetail_nmpoin"];
                        sumwin_1_listpatterndetail = sumwin_1_listpatterndetail + (minbet_listpatterndetail * record[i]["listpatterndetail_poin"]);
                        sumwin_2_listpatterndetail = sumwin_2_listpatterndetail + (minbet_listpatterndetail * record[i]["listpatterndetail_poin"])*2;
                        sumwin_3_listpatterndetail = sumwin_3_listpatterndetail + (minbet_listpatterndetail * record[i]["listpatterndetail_poin"])*3;
                        sumwin_4_listpatterndetail = sumwin_4_listpatterndetail + (minbet_listpatterndetail * record[i]["listpatterndetail_poin"])*4;
                    }else{
                        nm_poin = "";
                        
                    }
                    if(record[i]["listpatterndetail_poin"] > 0){
                        css_poin = "color:blue;";
                    }else{
                        css_poin = "color:red;";
                    }
                    sumbet_listpatterndetail = sumbet_listpatterndetail + minbet_listpatterndetail
                    listpatterndetail = [
                        ...listpatterndetail,
                        {
                            listpatterndetail_no: no,
                            listpatterndetail_id: record[i]["listpatterndetail_id"],
                            listpatterndetail_idpoin: record[i]["listpatterndetail_idpoin"],
                            listpatterndetail_nmpoin: nm_poin,
                            listpatterndetail_poin: record[i]["listpatterndetail_poin"],
                            listpatterndetail_poin_css: css_poin,
                            listpatterndetail_status: record[i]["listpatterndetail_status"],
                            listpatterndetail_status_css: record[i]["listpatterndetail_status_css"],
                        },
                    ];
                }
                sumwin_1_listpatterndetail = sumbet_listpatterndetail - sumwin_1_listpatterndetail
                sumwin_2_listpatterndetail = sumbet_listpatterndetail - sumwin_2_listpatterndetail
                sumwin_3_listpatterndetail = sumbet_listpatterndetail - sumwin_3_listpatterndetail
                sumwin_4_listpatterndetail = sumbet_listpatterndetail - sumwin_4_listpatterndetail
                if(sumwin_1_listpatterndetail>0){
                    sumwin_1_listpatterndetail_css = "color:blue;";
                }else{
                    sumwin_1_listpatterndetail_css = "color:red;";
                }
                if(sumwin_2_listpatterndetail>0){
                    sumwin_2_listpatterndetail_css = "color:blue;";
                }else{
                    sumwin_2_listpatterndetail_css = "color:red;";
                }
                if(sumwin_3_listpatterndetail>0){
                    sumwin_3_listpatterndetail_css = "color:blue;";
                }else{
                    sumwin_3_listpatterndetail_css = "color:red;";
                }
                if(sumwin_4_listpatterndetail>0){
                    sumwin_4_listpatterndetail_css = "color:blue;";
                }else{
                    sumwin_4_listpatterndetail_css = "color:red;";
                }
            }
        } else {
            logout();
        }
    }
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The Code is required\n"
            }
            
            if(name_field == ""){
                flag = false
                msg += "The Domain is required\n"
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
            const res = await fetch("/api/listpatternsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    listpattern_search: "",
                    listpattern_search_status: "",
                    listpattern_page: 0,
                    listpattern_id: idrecord.toUpperCase(),
                    listpattern_nmlistpattern: name_field,
                    listpattern_status: status_field,
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
    async function handleListDetailPatternSave() {
        let flag = true
        let msg = ""
        if(sDataListPatternDetail == "New"){
            if(idlistpattern_listpatterndetail_field == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(status_listpatterndetail_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idlistpattern_listpatterndetail_field == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(status_listpatterndetail_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/listpatterndetailsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataListPatternDetail,
                    page:"CURR-SAVE",
                    listpatterndetail_idlistpattern: idlistpattern_listpatterndetail_field,
                    listpatterndetail_status: status_listpatterndetail_field,
                    listpatterndetail_idpoin: parseInt(idpoin_listpatterndetail_field),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataListPatternDetail=="New"){
                    clearField_listpatterndetail()
                }
                msgloader = json.message;
                call_listpatterndetail(idlistpattern_listpatterndetail_field)
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
    async function handleListDetailPatternDeleteSingle(e,tipe) {
        let flag = true
        let msg = ""
        if(tipe == "SINGLE"){
            if(e == ""){
                flag = false
                msg += "The ID is required\n"
            }
        }
        
        if(tipe == ""){
            flag = false
            msg += "The Tipe is required\n"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/listpatterndetaildelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataListPatternDetail,
                    page:"CURR-SAVE",
                    listpatterndetail_id: parseInt(e),
                    Listpatterndetail_idlistpattern: idlistpattern_listpatterndetail_field,
                    listpatterndetail_tipe: tipe,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                msgloader = json.message;
                call_listpatterndetail(idlistpattern_listpatterndetail_field)
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }

    function clearField(){
        idrecord = "";
        name_field = "";
        status_field = "";
        flag_code = false
        create_field = "";
        update_field = "";
    }
    function clearField_listpatterndetail(){
        listpoint = []
        status_listpatterndetail_field = "N"
        flag_idpoin_field = false
        idpoin_listpatterndetail_field = "0"
        create_listpatterndetail_field = "";
        update_listpatterndetail_field = "";
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
    function statuswinlose(e){
        let result = "LOSE"
        if(e == "Y"){
            result = "WIN"
        }
        return result
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
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
                            placeholder="Search Listpattern"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CODE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                                <th NOWRAP width="35%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">POIN</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTAL_LOSE</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTAL_WIN</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                NewData("Edit",rec.home_id,rec.home_name, rec.home_status,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                Detaillistpattern(rec.home_id);
                                            }} class="bi bi-file-earmark-plus"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmpoin}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.home_totallose_css}">{rec.home_totallose}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.home_totalwin_css}">{rec.home_totalwin}</td>
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
	modal_size="modal-dialog-centered"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Code</label>
            <input bind:value={idrecord}
                use:uperCase  
                disabled={flag_code}
                class="required form-control"
                maxlength="30"
                type="text"
                placeholder="CODE"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <Input bind:value={name_field}
                class="required"
                type="text"
                placeholder="Name"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
            <select
                class="form-control required"
                bind:value={status_field} >
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
	modal_id="modallistpattern"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="List Pattern Detail - {idlistpattern_listpatterndetail_field}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
       <div class="padding:5px;">
            <table style="width: 100%;padding:5px;">
                <tr>
                    <td width="20%" NOWRAP style="font-size: 12px;">TOTAL LOSE</td>
                    <td width="1%" style="font-size: 12px;">:</td>
                    <td width="*" style="text-align:right;font-size: 12px;padding-right:5px;{totallose_listpatterndetail_css}">{totallose_listpatterndetail}</td>
                </tr>
                <tr>
                    <td style="font-size: 12px;">TOTAL WIN</td>
                    <td style="font-size: 12px;">:</td>
                    <td style="text-align:right;font-size: 12px;padding-right:5px;{totalwin_listpatterndetail_css}">{totalwin_listpatterndetail}</td>
                </tr>
                <tr>
                    <td style="font-size: 12px;">SUM BET</td>
                    <td style="font-size: 12px;">:</td>
                    <td style="text-align:right;font-size: 12px;color:red;padding-right:5px;">{new Intl.NumberFormat().format(sumbet_listpatterndetail)}</td>
                </tr>
                <tr>
                    <td style="font-size: 12px;">WINLOSE ROUND BET 1</td>
                    <td style="font-size: 12px;">:</td>
                    <td style="text-align:right;font-size: 12px;padding-right:5px;{sumwin_1_listpatterndetail_css}">{new Intl.NumberFormat().format(sumwin_1_listpatterndetail)}</td>
                </tr>
                <tr>
                    <td style="font-size: 12px;">WINLOSE ROUND BET 2</td>
                    <td style="font-size: 12px;">:</td>
                    <td style="text-align:right;font-size: 12px;padding-right:5px;{sumwin_2_listpatterndetail_css}">{new Intl.NumberFormat().format(sumwin_2_listpatterndetail)}</td>
                </tr>
                <tr>
                    <td style="font-size: 12px;">WINLOSE ROUND BET 3</td>
                    <td style="font-size: 12px;">:</td>
                    <td style="text-align:right;font-size: 12px;padding-right:5px;{sumwin_3_listpatterndetail_css}">{new Intl.NumberFormat().format(sumwin_3_listpatterndetail)}</td>
                </tr>
                <tr>
                    <td style="font-size: 12px;">WINLOSE ROUND BET 4</td>
                    <td style="font-size: 12px;">:</td>
                    <td style="text-align:right;font-size: 12px;padding-right:5px;{sumwin_4_listpatterndetail_css}">{new Intl.NumberFormat().format(sumwin_4_listpatterndetail)}</td>
                </tr>
            </table>
       </div>
    </slot:template>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">WIN</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">POIN</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BET</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ROUND 1</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ROUND 2</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ROUND 3</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ROUND 4</th>
                </tr>
            </thead>
            <tbody>
                {#each listpatterndetail as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font}; cursor:pointer;">
                            <i on:click={() => {
                                Deletelistpatterndetail(rec.listpatterndetail_id,"SINGLE");
                            }} class="bi bi-trash"></i>
                        </td>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.listpatterndetail_no}</td>
                        <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                            <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.listpatterndetail_status_css}">
                                {statuswinlose(rec.listpatterndetail_status)}
                            </span>
                            
                        </td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.listpatterndetail_nmpoin}</td>
                        <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.listpatterndetail_poin_css}">{rec.listpatterndetail_poin}x</td>
                        <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">
                            {new Intl.NumberFormat().format(parseInt(minbet_listpatterndetail))}
                        </td>
                        <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.listpatterndetail_poin_css}">
                            {new Intl.NumberFormat().format(parseInt(rec.listpatterndetail_poin) * parseInt(minbet_listpatterndetail))}
                        </td>
                        <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.listpatterndetail_poin_css}">
                            {new Intl.NumberFormat().format((parseInt(rec.listpatterndetail_poin) * parseInt(minbet_listpatterndetail))*2)}
                        </td>
                        <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.listpatterndetail_poin_css}">
                            {new Intl.NumberFormat().format((parseInt(rec.listpatterndetail_poin) * parseInt(minbet_listpatterndetail))*3)}
                        </td>
                        <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.listpatterndetail_poin_css}">
                            {new Intl.NumberFormat().format((parseInt(rec.listpatterndetail_poin) * parseInt(minbet_listpatterndetail))*4)}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
            call_formlistpatterndetail();
        }} 
        button_title="New"
        button_css="btn-info"/>

        <Button on:click={() => {
            Deletelistpatterndetail(0,"ALL");
        }} 
        button_title="Clear All"
        button_css="btn-danger"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalcrudlistdetailcompany"
	modal_size="modal-dialog-centered"
	modal_title=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status Card</label>
            <select
                class="form-control required"
                on:change={handleSelectStatus}
                bind:value={status_listpatterndetail_field}>
                <option value="N">LOSE</option>
                <option value="Y">WIN</option>
            </select>
        </div>
        {#if flag_idpoin_field}
        <div class="mb-3">
            <label for="exampleForm" class="form-label">List Poin</label>
            <select
                class="form-control required"
                bind:value={idpoin_listpatterndetail_field}>
                {#each listpoint as rec}
                    <option value="{rec.lispoint_id}">{rec.lispoint_name}</option>
                {/each}
            </select>
        </div>
        {/if}
        {#if sDataListPatternDetail != "New"}
            <div class="mb-3">
                <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                    Create : {create_listpatterndetail_field}<br />
                    Update : {update_listpatterndetail_field}
                </div>
            </div>
        {/if}
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleListDetailPatternSave();
            }} 
            
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>