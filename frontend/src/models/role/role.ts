import { AuthUser } from "../../userContext";

export const CHIEF = "チーフ";
export const INFRA = "インフラ";
export const isChief = (authUser?: AuthUser) => {
  return authUser && (authUser.role_list?.includes(CHIEF) ?? []);
};
export const isInfra = (authUser?: AuthUser) => {
  return authUser && (authUser.role_list?.includes(INFRA) ?? []);
};
