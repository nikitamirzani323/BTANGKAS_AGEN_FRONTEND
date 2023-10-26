<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

	export let table_header_font
	export let table_body_font
	export let token
	export let listAdmin = []
	export let listAdminrule = []
	export let totalrecord = 0
	export let company_db = ""
    let dispatch = createEventDispatcher();
	let title_page = "Admin"
    let sData = "";
    let username_flag = false;
    let company_flag = false;
    let adminrule_flag = false;
    let myModal_newentry = ""
    let css_loader = "display: none;";
    let msgloader = "";

    let idrecord = ""
    let flag_btnsave = true;
    let listcompanyadminrule = []
    let company_field = ""
    let rule_field = ""
    let username_field = ""
    let password_field = ""
    let name_field = ""
    let phone1_field = ""
    let phone2_field = ""
    let status_field = ""
    let create_field = ""
    let update_field = ""


    const NewData = (e,idrec,idcompany,idrule,username,name,phone1,phone2,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            username_flag = true
            company_flag = true
            adminrule_flag = true
            idrecord = idrec
            company_field = idcompany
            rule_field = idrule
            username_field = username
            name_field = name
            phone1_field = phone1
            phone2_field = phone2
            status_field = status
            create_field = create
            update_field = update

            listcompanyadminrule = [];
            for(let i=0;i<listAdminrule.length;i++){
                if(listAdminrule[i].companyadminrule_idcompany == idcompany){
                    listcompanyadminrule = [...listcompanyadminrule,
                        {
                        companyadminrule_id: listAdminrule[i]["companyadminrule_id"],
                        companyadminrule_idcompany: listAdminrule[i]["companyadminrule_idcompany"],
                        companyadminrule_nmrule: listAdminrule[i]["companyadminrule_nmrule"],
                        },
                    ];
                }
            }
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(company_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(rule_field == ""){
                flag = false
                msg += "The Rule is required\n"
            }
            if(username_field == ""){
                flag = false
                msg += "The Username is required\n"
            }
            if(password_field == ""){
                flag = false
                msg += "The Password is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone1 is required\n"
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
            if(company_field == ""){
                flag = false
                msg += "The Company is required\n"
            }
            if(rule_field == ""){
                flag = false
                msg += "The Rule is required\n"
            }
            if(username_field == ""){
                flag = false
                msg += "The Username is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone1 is required\n"
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
            const res = await fetch("/api/saveadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    admin_id: idrecord,
                    admin_idrule: parseInt(rule_field),
                    admin_username: company_field.toLowerCase()+username_field.toLowerCase(),
                    admin_password: password_field,
                    admin_nama: name_field,
                    admin_phone1: phone1_field,
                    admin_phone2: phone2_field,
                    admin_status: status_field,
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
        rule_field = ""
        username_field = ""
        password_field = ""
        name_field = ""
        phone1_field = ""
        phone2_field = ""
        status_field = ""
        create_field = ""
        update_field = ""
        company_flag = false;
        username_flag = false;
        adminrule_flag = false;
    }
    $:{
    //    console.log("nama : "+company_db)
       company_field = company_db
    }
    
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    function lowerCase(element) {
        function onInput(event) {
            element.value = element.value.toLowerCase();
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
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-body">
                        <table class="table table-striped ">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">USERNAME</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">RULE</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NAME</th>
                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PHONE</th>
                                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">LAST LOGIN</th>
                                    <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">LAST IPADDRESS</th>
                                </tr>
                            </thead>
                            {#if totalrecord > 0}
                            <tbody>
                                {#each listAdmin as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            {#if rec.admin_rule != "MASTER"}
                                            <i 
                                                on:click={() => {
                                                    //e,idcompany,idrule,username,name,phone1,phone2,status
                                                    NewData("Edit",rec.admin_id,rec.admin_idcompany,rec.admin_idrule,
                                                    rec.admin_username,rec.admin_nama,rec.admin_phone1,rec.admin_phone2,
                                                    rec.admin_status,rec.admin_create,rec.admin_update);
                                                }} 
                                                class="bi bi-pencil"></i>
                                            {/if}
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                            <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.admin_statuscss}">
                                                {status(rec.admin_status)}
                                            </span>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_username}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_rule}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_nama}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.admin_phone1} / {rec.admin_phone2}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_lastlogin}</td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.admin_lastipaddres}</td>
                                    </tr>
                                {/each}
                            </tbody>
                            {:else}
                            <tbody>
                                <tr>
                                    <td colspan="10">
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
	modal_id="modalentry"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Rule</label>
                    <select
                        bind:value={rule_field} 
                        disabled='{adminrule_flag}'
                        class="form-control required">
                        <option value="">--Select--</option>
                        {#each listAdminrule as rec }
                            <option value="{rec.adminrule_id}">{rec.adminrule_name}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Username</label>
                    <div class="input-group mb-3">
                        {#if sData == "New"}
                        <span class="input-group-text" id="basic-addon1">{company_db.toLowerCase()}</span>
                        {/if}
                        <input
                            use:lowerCase 
                            bind:value={username_field}
                            disabled='{username_flag}'
                            type="text"
                            class="form-control required"
                            maxlength="10"
                            placeholder="Username"
                            aria-label="Username"/>
                    </div>
                    
                   
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Password</label>
                    <input
                        bind:value={password_field}
                        type="password"
                        class="form-control "
                        placeholder="Password"
                        aria-label="Password"/>
                </div>
            </div>
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <input
                        bind:value={name_field}
                        type="text"
                        class="form-control required"
                        placeholder="Name"
                        aria-label="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 1</label>
                    <input
                        bind:value={phone1_field}
                        type="text"
                        class="form-control required"
                        placeholder="Phone1"
                        aria-label="Phone1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 2</label>
                    <input
                        bind:value={phone2_field}
                        type="text"
                        class="form-control "
                        placeholder="Phone2"
                        aria-label="Phone2"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="">--Select--</option>
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