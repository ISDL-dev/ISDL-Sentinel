import React, { useState, useEffect, ChangeEvent, useRef } from "react";
import {
  Box,
  Flex,
  Grid,
  Select,
  Text,
  Button,
  chakra,
  useToast,
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel,
  Avatar,
  TableContainer,
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
} from "@chakra-ui/react";
import dayjs from "dayjs";
import "dayjs/locale/ja";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import html2canvas from "html2canvas";
import jsPDF from "jspdf";
import { labAssistantApi } from "../../api";
import {
  GetLabAssistantMember200ResponseInner,
  GetLabAssistantSchedule200ResponseInner,
  PostLabAssistantScheduleRequestInner,
} from "../../schema";
import { useNavigate } from "react-router-dom";

dayjs.locale("ja");

interface Shift {
  date: number;
  user_name: string;
}

const generateDaysInMonth = (year: number, month: number): Shift[] => {
  const startOfMonth = dayjs().year(year).month(month).startOf("month");
  const endOfMonth = dayjs().year(year).month(month).endOf("month");
  const daysInMonth: Shift[] = [];

  for (let i = 0; i < startOfMonth.day(); i++) {
    daysInMonth.push({
      date: 0,
      user_name: "",
    });
  }

  for (
    let date = startOfMonth;
    date.isBefore(endOfMonth) || date.isSame(endOfMonth);
    date = date.add(1, "day")
  ) {
    daysInMonth.push({
      date: date.date(),
      user_name: "",
    });
  }
  return daysInMonth;
};

const decodeDate = (dateString: string) => {
  const date = dayjs(dateString);
  return `${dayjs(date).format("YYYY年MM月DD日")}（${dayjs(date).format(
    "ddd"
  )})`;
};

const formatMonth = (date: Date) => {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  return `${year}-${month}`;
};

const fetchLabAssistantData = async (
  selectedDate: Date | null,
  selectedYear: number,
  selectedMonth: number,
  setLabAssistantMember: React.Dispatch<React.SetStateAction<GetLabAssistantMember200ResponseInner[]>>,
  setShifts: React.Dispatch<React.SetStateAction<Shift[]>>,
  labAssistantApi: any
) => {
  const memberResponse = await labAssistantApi.getLabAssistantMember();
  setLabAssistantMember(memberResponse.data);

  const formattedMonth = formatMonth(selectedDate || new Date());
  const scheduleResponse = await labAssistantApi.getLabAssistantSchedule(formattedMonth);

  const updatedShifts = generateDaysInMonth(
    selectedYear,
    selectedMonth
  ).map((shift) => {
    const schedule = scheduleResponse.data?.find(
      (s: GetLabAssistantSchedule200ResponseInner) =>
        dayjs(s.shift_date).date() === shift.date
    );
    return schedule ? { ...shift, user_name: schedule.user_name } : shift;
  });

  setShifts(updatedShifts);
};

export default function Profile() {
  const [selectedYear, setSelectedYear] = useState<number>(dayjs().year());
  const [selectedMonth, setSelectedMonth] = useState<number>(dayjs().month());
  const [shifts, setShifts] = useState<Shift[]>(
    generateDaysInMonth(selectedYear, selectedMonth)
  );
  const [labAssistantMember, setLabAssistantMember] = useState<
    GetLabAssistantMember200ResponseInner[]
  >([]);
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const toast = useToast();
  const navigate = useNavigate();
  const gridRef = useRef<HTMLDivElement>(null); 
  const titleRef = useRef<HTMLDivElement>(null); 

  useEffect(() => {
    fetchLabAssistantData(
      selectedDate,
      selectedYear,
      selectedMonth,
      setLabAssistantMember,
      setShifts,
      labAssistantApi
    );
  }, [selectedYear, selectedMonth]);

  const handleMonthChange = (date: Date | null) => {
    if (date) {
      setSelectedYear(dayjs(date).year());
      setSelectedMonth(dayjs(date).month());
      setShifts(generateDaysInMonth(dayjs(date).year(), dayjs(date).month()));
      setSelectedDate(date);
    }
  };

  const handleNameChange = (
    index: number,
    event: ChangeEvent<HTMLSelectElement>
  ) => {
    const newShifts = [...shifts];
    newShifts[index].user_name = event.target.value;
    setShifts(newShifts);
  };

  const handleSubmit = async () => {
    try {
      const requestBody: PostLabAssistantScheduleRequestInner[] = shifts
        .filter((shift) => shift.date !== 0 && shift.user_name !== "")
        .map((shift) => {
          const member = labAssistantMember.find(
            (m) => m.user_name === shift.user_name
          );
          return {
            user_id: member?.user_id || 0,
            shift_date: dayjs(
              `${selectedYear}-${selectedMonth + 1}-${shift.date}`
            ).format("YYYY-MM-DD"),
          };
        });

      const formattedMonth = formatMonth(selectedDate || new Date());
      await labAssistantApi.postLabAssistantSchedule(
        formattedMonth,
        requestBody
      );

      await fetchLabAssistantData(
        selectedDate,
        selectedYear,
        selectedMonth,
        setLabAssistantMember,
        setShifts,
        labAssistantApi
      );
      
      toast({
        title: "シフトが正常に登録されました。",
        status: "success",
        duration: 5000,
        isClosable: true,
      });
    } catch (err) {
      toast({
        title: "シフトの登録に失敗しました。",
        status: "error",
        duration: 5000,
        isClosable: true,
      });
    }
  };

  const handlePdfDownload = async () => {
    const titleElement = titleRef.current;
    const gridElement = gridRef.current;
  
    if (!titleElement || !gridElement) {
      console.error('Element not found');
      return;
    }
  
    const titleCanvas = await html2canvas(titleElement, { scale: 2 });
  
    const gridCanvas = await html2canvas(gridElement, { scale: 2 });
  
    const titleImgData = titleCanvas.toDataURL('image/png');
    const gridImgData = gridCanvas.toDataURL('image/png');
    const pdf = new jsPDF('landscape');
  
    const titleFontSize = 16;
    pdf.setFontSize(titleFontSize);
    
    const titleWidth = 50;
    const titleHeight = 10;
    
    const titleX = (pdf.internal.pageSize.width - 210) / 2;
    const titleY = 40;
  
    pdf.addImage(titleImgData, 'PNG', titleX, titleY, titleWidth, titleHeight);
  
    const imgWidth = 210; 
    const imgHeight = gridCanvas.height * imgWidth / gridCanvas.width;
  
    const gridX = (pdf.internal.pageSize.width - imgWidth) / 2;
    const gridY = titleY + titleHeight;
  
    pdf.addImage(gridImgData, 'PNG', gridX, gridY, imgWidth, imgHeight);
  
    const filename = `${selectedYear}${String(selectedMonth + 1).padStart(2, '0')}.pdf`;
    pdf.save(filename);
  };
  
  

  return (
    <Box>
      <Tabs>
        <TabList>
          <Tab>シフトボード</Tab>
          <Tab>シフトメンバー</Tab>
        </TabList>

        <TabPanels>
          <TabPanel>
            {/* シフトボード */}
            <Flex
              flexDirection="row"
              mb={3}
              alignItems="center"
              justifyContent="space-between"
            >
              <Box ref={titleRef}>
                <Text fontSize="xl" fontWeight="bold" mb={3} textAlign="left">
                  {`${selectedYear}年${selectedMonth + 1}月LAシフト表`}
                </Text>
              </Box>
              <Flex alignItems="center" gap={4}>
                <chakra.label htmlFor="month-picker" fontWeight="bold">
                  月を選択:
                </chakra.label>
                <DatePicker
                  selected={selectedDate}
                  onChange={handleMonthChange}
                  dateFormat="yyyy/MM"
                  showMonthYearPicker
                  id="month-picker"
                  className="custom-datepicker"
                />
                <Button colorScheme="blue" onClick={handlePdfDownload}>
                  PDF化
                </Button>
                <Button colorScheme="teal" onClick={handleSubmit}>
                  登録
                </Button>
              </Flex>
            </Flex>

            <Box ref={gridRef}>
              <Grid templateColumns="repeat(7, 1fr)" gap={1} mt={4}>
                {["日", "月", "火", "水", "木", "金", "土"].map(
                  (day, index) => (
                    <Box
                      key={index}
                      p={2}
                      bg="rgba(79, 209, 197, 1)"
                      textAlign="center"
                    >
                      {day}
                    </Box>
                  )
                )}
                {shifts.map((shift, index) => (
                  <Box
                    key={index}
                    p={2}
                    border="1px solid"
                    borderColor="gray.200"
                    bg={shift.date === 0 ? "transparent" : "white"}
                  >
                    {shift.date !== 0 && (
                      <>
                        <Text mb={2} textAlign="center">
                          {shift.date}日
                        </Text>
                        <Select
                          value={shift.user_name}
                          onChange={(event) => handleNameChange(index, event)}
                          placeholder=" "
                        >
                          {labAssistantMember.map((member) => (
                            <option
                              key={member.user_id}
                              value={member.user_name}
                            >
                              {member.user_name}
                            </option>
                          ))}
                        </Select>
                      </>
                    )}
                  </Box>
                ))}
              </Grid>
            </Box>
          </TabPanel>

          <TabPanel>
            <TableContainer
              pb={14}
              pr={{ base: 2, md: 14 }}
              pl={{ base: 2, md: 14 }}
              mt={8}
              outlineOffset={2}
              overflowX="unset"
              overflowY="scroll"
              height="65vh"
            >
              <Table
                size="lg"
                border="2px"
                borderColor="gray.200"
                variant="simple"
              >
                <Thead>
                  <Tr bgColor="#E6EBED">
                    <Th>ユーザー</Th>
                    <Th>最終LA日程</Th>
                    <Th>LA回数</Th>
                  </Tr>
                </Thead>
                <Tbody outline="1px">
                  {labAssistantMember.map((member) => (
                    <Tr key={member.user_id}>
                      <Td>
                        <Flex alignItems="center" gap={3}>
                          <Avatar
                            size={"md"}
                            src={`./avatar/${member.avatar_img_path}`}
                            border="2px"
                            onClick={() =>
                              navigate("/profile", {
                                state: { userId: member.user_id },
                              })
                            }
                          />
                          {member.user_name}
                        </Flex>
                      </Td>
                      <Td>{decodeDate(member.last_shift_date)}</Td>
                      <Td>{member.count}</Td>
                    </Tr>
                  ))}
                </Tbody>
              </Table>
            </TableContainer>
          </TabPanel>
        </TabPanels>
      </Tabs>
    </Box>
  );
}
