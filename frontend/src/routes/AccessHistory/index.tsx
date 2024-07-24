import { useEffect, useState } from "react";
import {
  Grid,
  Table,
  TableContainer,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  Flex,
  Avatar,
  chakra,
  Box,
  Text,
} from "@chakra-ui/react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "./access_history.css";
import { accessHistoryApi } from "../../api";
import { GetAccessHistory200ResponseInner } from "../../schema";

export default function AccessHistory() {
  const [accessHistory, setAccessHistoryData] = useState<GetAccessHistory200ResponseInner[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());

  // Handle date change from DatePicker
  const handleDateChange = (date: Date | null) => {
    if (date) {
      setSelectedDate(date);
    }
  };

  // Format date as 'yyyy/MM'
  const formatDate = (date: Date) => {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    return `${year}-${month}`;
  };

  // Extract time from date string in '2024-07-01T09:00:00Z' format
  const formatTime = (dateTime: string) => {
    const date = new Date(dateTime);
    return date.toISOString().substring(11, 19);
  };

  // Fetch data based on the selected date
  useEffect(() => {
    const fetchUserData = async (date: Date) => {
      setLoading(true); // Set loading to true when fetching data
      try {
        const formattedDate = formatDate(date);
        const response = await accessHistoryApi.getAccessHistory(formattedDate);
        setAccessHistoryData(response.data ?? []);
      } catch (err) {
        setError('データの取得に失敗しました');
      } finally {
        setLoading(false); // Set loading to false when done
      }
    };

    // Fetch data with initial date
    fetchUserData(selectedDate);
  }, [selectedDate]); // Dependency array includes selectedDate

  // Initialize selectedDate to the current month
  useEffect(() => {
    const now = new Date();
    // Set to the first day of the current month
    const currentMonth = new Date(now.getFullYear(), now.getMonth(), 1);
    setSelectedDate(currentMonth);
  }, []); // Empty dependency array ensures this runs only on initial render

  if (loading) return <p>Loading...</p>;
  if (error) return <p>{error}</p>;

  return (
    <Box p={6} mt={{ base: 20, md: 0 }}>
      <Grid
        templateColumns={{ base: "1fr", md: "1fr auto" }}
        alignItems="center"
      >
        <Text
          fontSize={{ base: "2xl", md: "4xl" }}
          textAlign="left"
          className="block font-bold text-gray-900 dark:text-white"
        >
          入退室履歴
        </Text>
        <Flex
          justifyContent={{ base: "center", md: "flex-end" }}
          alignItems="center"
          mt={{ base: 4, md: 0 }}
          flexDirection={{ base: "column", md: "row" }}
          textAlign={{ base: "center", md: "left" }}
        >
          <chakra.label htmlFor="month-picker" mr={3} fontWeight="bold">
            月を選択:
          </chakra.label>
          <DatePicker
            selected={selectedDate}
            onChange={handleDateChange}
            dateFormat="yyyy/MM"
            showMonthYearPicker
            id="month-picker"
            className="custom-datepicker"
          />
        </Flex>
      </Grid>
      <TableContainer
        pb={14}
        pr={{ base: 2, md: 14 }} // Reduced padding on mobile
        pl={{ base: 2, md: 14 }}
        mt={8}
        outlineOffset={2}
        overflowX="unset"
        overflowY="scroll"
        height="65vh"
      >
        <Table size="lg" border="2px" borderColor="gray.200" variant="simple">
          <Thead top={0}>
            <Tr bgColor="#E6EBED">
              <Th w="20%">日付</Th>
              <Th w="20%">入室者</Th>
              <Th w="20%">入室時間</Th>
              <Th w="20%">退室者</Th>
              <Th w="20%">退室時間</Th>
            </Tr>
          </Thead>
          <Tbody outline="1px">
            {accessHistory.length === 0 ? (
              <Tr>
                <Td colSpan={5} textAlign="center">
                  データがありません
                </Td>
              </Tr>
            ) : (
              accessHistory.map((access) => (
                <Tr key={access.date}>
                  <Td>{access.date}</Td>
                  <Td>
                    <Flex alignItems={"center"} gap={3}>
                      <Avatar
                        size={"md"}
                        src={`./avatar/${access.entering.avatar_img_path}`}
                        border="2px"
                      />
                      {access.entering.user_name}
                    </Flex>
                  </Td>
                  <Td>{formatTime(access.entering.entered_at)}</Td>
                  <Td>
                    <Flex alignItems={"center"} gap={3}>
                      <Avatar
                        size={"md"}
                        src={`./avatar/${access.leaving.avatar_img_path}`}
                        border="2px"
                      />
                      {access.leaving.user_name}
                    </Flex>
                  </Td>
                  <Td>{formatTime(access.leaving.left_at)}</Td>
                </Tr>
              ))
            )}
          </Tbody>
        </Table>
      </TableContainer>
    </Box>
  );
}
