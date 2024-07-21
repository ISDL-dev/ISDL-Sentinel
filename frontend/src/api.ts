import axios from "axios";
import { 
  GetAttendeesListApi, 
  PutStatusApi, 
  GetUserByIdApi, 
  PutAvatarApi 
} from "./schema";

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

const api = axios.create({
  baseURL: baseURL,
});
const getAttendeesListApi = new GetAttendeesListApi(undefined, baseURL, api);
const putStatusApi = new PutStatusApi(undefined, baseURL, api);
const getUserByIdApi = new GetUserByIdApi(undefined, baseURL, api);
const putAvatarApi = new PutAvatarApi(undefined, baseURL, api);

export {
  getAttendeesListApi,
  putStatusApi,
  getUserByIdApi,
  putAvatarApi
}
