import {defineStore} from "pinia";
import {reactive, ref} from "vue";

export const useFileFolderStore = defineStore('file-folder', () => {
    const selectedItems: { files: number[], folders: number[] } = reactive({
        files: [],
        folders: []
    })

    let folderId = ref<number>(0)

    function selectChange(ids: number[], forFile: boolean) {
        if (forFile) {
            selectedItems.files = ids
        } else {
            selectedItems.folders = ids
        }
    }

    function setFolderId(id: number) {
        folderId.value = id
    }

    return {
        selectedItems, folderId,
        selectChange, setFolderId
    }
})