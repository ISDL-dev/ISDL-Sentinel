import { GetUsersInfo200ResponseInner } from "../../schema";

export const gradeOrder = ["Teacher", "D3", "D2", "D1", "M2", "M1", "U4"];
export const sortUsersByGrade = (
  users: GetUsersInfo200ResponseInner[]
): GetUsersInfo200ResponseInner[] => {
  return users.sort((a, b) => {
    if (a.grade === "OB" && b.grade !== "OB") return 1;
    if (a.grade !== "OB" && b.grade === "OB") return -1;

    if (a.grade === b.grade) {
      return a.user_id - b.user_id;
    }
    return gradeOrder.indexOf(a.grade) - gradeOrder.indexOf(b.grade);
  });
};
