import { GetUsersInfo200ResponseInner } from "../../schema";

export const GradeList = {
  Teacher: "Teacher",
  D3: "D3",
  D2: "D2",
  D1: "D1",
  M2: "M2",
  M1: "M1",
  U4: "U4",
  OB: "OB",
};

export const gradeOrder = [
  GradeList.Teacher,
  GradeList.D3,
  GradeList.D2,
  GradeList.D1,
  GradeList.M2,
  GradeList.M1,
  GradeList.U4,
];

export const sortUsersByGrade = (
  users: GetUsersInfo200ResponseInner[]
): GetUsersInfo200ResponseInner[] => {
  return users.sort((a, b) => {
    if (a.grade === GradeList.OB && b.grade !== GradeList.OB) return 1;
    if (a.grade !== GradeList.OB && b.grade === GradeList.OB) return -1;

    if (a.grade === b.grade) {
      return a.user_id - b.user_id;
    }
    return gradeOrder.indexOf(a.grade) - gradeOrder.indexOf(b.grade);
  });
};
