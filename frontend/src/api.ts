import axios from "axios";
import { GetAttendeesListApi, PutStatusApi } from "./schema";

const baseURL = process.env.REACT_APP_BACKEND_ENDPOINT;

const api = axios.create({
  baseURL: baseURL,
});

export const putStatusApi = new PutStatusApi(undefined, baseURL, api);
export const getAttendeesListApi = new GetAttendeesListApi(
  undefined,
  baseURL,
  api
);
