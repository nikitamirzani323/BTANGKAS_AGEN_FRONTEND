<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import Button from "../../components/Button.svelte";
    import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    export let table_header_font;
    export let table_body_font;
    export let token;
    export let listAdminrule = [];
    export let totalrecord = 0;
    let dispatch = createEventDispatcher();
    let title_page = "Admin Rule";
    let sData = "";
    let myModal_newentry = "";
    let css_loader = "display: none;";
    let msgloader = "";
    let flag_btnsave = true;
    let idrecord = 0;
    let name_field = "";
    let rule_field = "";
    
    const NewData = () => {
        clearField();
        sData = "New";
        myModal_newentry = new bootstrap.Modal(
            document.getElementById("modalentry")
        );
        myModal_newentry.show();
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const EditData = (a,b,c) => {
        const adminrule = {
            a,
            b,
            c,
        };
        dispatch("handleEditData", adminrule);
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/saveadminrule", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    adminrule_id: parseInt(idrecord),
                    adminrule_name:name_field.toUpperCase(),
                    adminrule_rule: "",
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
    function clearField() {
        
    }
   

    function callFunction(event) {
        switch (event.detail) {
            case "NEW":
                NewData();
                break;
            case "REFRESH":
                RefreshHalaman();
                break;
            case "SAVE":
                handleSubmit();
                break;
        }
    }
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
            <Panel card_title={title_page} card_footer={totalrecord}>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">RULE</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                            <tbody>
                                {#each listAdminrule as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align: top;cursor:pointer;">
                                            {#if rec.adminrule_name != "MASTER"}
                                            <i on:click={() => {
                                                    EditData(
                                                        rec.adminrule_id,
                                                        rec.adminrule_name,
                                                        rec.adminrule_rule
                                                    );
                                                }} class="bi bi-pencil"/>
                                            {/if}
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.adminrule_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.adminrule_name}</td>
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
    modal_size="modal-dialog-centered"
    modal_title={title_page + "/" + sData}
    modal_footer_css="padding:5px;"
    modal_footer={true}>
    <slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <input
                bind:value={name_field}
                use:uperCase 
                type="text"
                class="form-control required"
                placeholder="Name"
                aria-label="Name"/>
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
