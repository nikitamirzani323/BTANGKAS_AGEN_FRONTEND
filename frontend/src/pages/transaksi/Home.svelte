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
	let title_page = "TRANSAKSI"
    let sData = "";
    let myModal_newentry = "";
    let flag_btnsave = true;
    let minbet_field = 0;
    let create_field = "";
    let update_field = "";
    let idrecord = 0;
 

      //COMPANY CONF POINT
    let sDataConfPoint = "";
    let listdetail_transaksi = [];
    let idrecord_confpoint_field = 0;
    let idbet_confpoint_field = 0;
    let point_confpoint_field = 0;
    let nmpoint_confpoint_field = "";
    let create_confpoint_field = "";
    let update_confpoint_field = "";


    let title_minbet = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
    }
    
    const NewData = (e,id) => {
        sData = e
        if(sData == "New"){
            
        }else{
            // flag_code = true;
            idrecord = id
            title_minbet = id;
            call_detail(id) 
            myModal_newentry = new bootstrap.Modal(document.getElementById("modallistconfpoint"));
            myModal_newentry.show();
        }
       
        
    };
    const EditConfPoint = (id,idbet,nmpoin,poin,create,update) => {
        sDataConfPoint = "Edit"
        idrecord_confpoint_field = id
        idbet_confpoint_field = idbet
        nmpoint_confpoint_field = nmpoin
        point_confpoint_field = poin
        create_confpoint_field = create;
        update_confpoint_field = update;
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalformconfpoint"));
        myModal_newentry.show();
        
    };
    async function handleSave_confpoint_generate() {
        let flag = true
        let msg = ""
        if(sDataConfPoint == "New"){
            
            if(idbet_confpoint_field == ""){
                flag = false
                msg += "The IDBET is required\n"
            }
        }else{
            if(idrecord_confpoint_field == ""){
                flag = false
                msg += "The ID is required\n"
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
            const res = await fetch("/api/listbetconfsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataConfPoint,
                    page:"CURR-SAVE",
                    conf_id: parseInt(idrecord_confpoint_field),
                    conf_idbet: parseInt(idbet_confpoint_field),
                    conf_minbet:parseInt(point_confpoint_field),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    // clearField_listbet()
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
    async function call_detail(idinvoice) {
        listdetail_transaksi = [];
        const res = await fetch("/api/transaksidetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                invoice: idinvoice,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listdetail_transaksi = [
                        ...listdetail_transaksi,
                        {
                        detail_no: no,
                        detail_id: record[i]["transaksidetail_id"],
                        detail_date: record[i]["transaksidetail_date"],
                        detail_round: record[i]["transaksidetail_roundbet"],
                        detail_bet: record[i]["transaksidetail_bet"],
                        detail_win: record[i]["transaksidetail_win"],
                        detail_bonus: record[i]["transaksidetail_bonus"],
                        detail_card_codepoin: record[i]["transaksidetail_card_codepoin"],
                        detail_card_win: record[i]["transaksidetail_card_win"],
                        detail_creditbefore: record[i]["transaksidetail_creditbefore"],
                        detail_creditafter: record[i]["transaksidetail_creditafter"],
                        detail_status: record[i]["transaksidetail_status"],
                        detail_create: record[i]["transaksidetail_create"],
                        detail_update: record[i]["transaksidetail_update"],
                        },
                    ];
                }
            }
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(minbet_field == ""){
                flag = false
                msg += "The Minbet is required\n"
            }
        }else{
            if(minbet_field == ""){
                flag = false
                msg += "The Minbet is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/listbetsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    lisbet_id: parseInt(idrecord),
                    lisbet_minbet: parseInt(minbet_field),
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
    
    function clearField(){
        idrecord = "";
        minbet_field = 0;
        flag_code = false
        create_field = "";
        update_field = "";
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
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
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
                            placeholder="Search Invoice"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">INVOICE</th>
                                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DATE</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">USERNAME</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PATTERN</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">POIN</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">ROUNDBET</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTALBET</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTALBONUS</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTALWIN</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">WINLOSE</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                NewData("Edit",rec.home_id);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_date}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_username}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_card_pattern}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_card_codepoin}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.home_roundbet)}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.home_totalbet)}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">{new Intl.NumberFormat().format(rec.home_totalbonus)}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">{new Intl.NumberFormat().format(rec.home_totalwin)}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.home_winlose_css}">{new Intl.NumberFormat().format(rec.home_winlose)}</td>
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
            <label for="exampleForm" class="form-label">Minbet</label>
            
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
	modal_id="modallistconfpoint"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="Invoice - {title_minbet}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DATE</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ROUND</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BET</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">WIN</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BONUS</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BEFORE</th>
                    <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">AFTER</th>
                </tr>
            </thead>
            <tbody>
                {#each listdetail_transaksi as rec }
                    <tr>
                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.detail_no}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.detail_status}</td>
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.detail_date}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_round)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_bet)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">{new Intl.NumberFormat().format(rec.detail_win)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:red;">{new Intl.NumberFormat().format(rec.detail_bonus)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_creditbefore)}</td>
                        <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_creditafter)}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
	</slot:template>
	<slot:template slot="footer">
        
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
