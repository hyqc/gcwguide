import axios from "@/util/request.js";

const fileApi = {
  add: "websites/add",
  list: "websites/list",
  delete: "websites/delete",
  update: "websites/update",
  groups: "websites/groups",
  imageUpload: "upload/image",
};

export function WebList(data) {
  return axios({
    method: "get",
    data: data,
    url: fileApi.list,
  });
}

export function WebAdd(data) {
  return axios({
    method: "post",
    data: data,
    url: fileApi.add,
  });
}


export function WebEdit(data) {
  return axios({
    method: "post",
    data: data,
    url: fileApi.update,
  });
}

export function WebDelete(data) {
  return axios({
    method: "post",
    data: data,
    url: fileApi.delete,
  });
}

export function WebGroups(data) {
  return axios({
    method: "get",
    data: data,
    url: fileApi.groups,
  });
}

export const WebLogoUploadUrl = fileApi.imageUpload
