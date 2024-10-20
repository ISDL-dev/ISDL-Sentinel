import {
  Flex,
  IconButton,
  Input,
  InputGroup,
  InputLeftAddon,
  InputRightAddon,
  Spacer,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  useColorModeValue,
} from "@chakra-ui/react";
import { KeyboardEvent, useEffect, useState } from "react";
import { RankingList } from "../../features/RankingList";
import {
  ChevronLeftIcon,
  ChevronRightIcon,
  TriangleDownIcon,
  TriangleUpIcon,
} from "@chakra-ui/icons";
import { rankingApi } from "../../api";
import { GetRanking200ResponseInner } from "../../schema";
import { Loading } from "../../features/Loading/Loading";

const getCurrentYearMonth = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = (now.getMonth() + 1).toString().padStart(2, "0");
  return `${year}-${month}`;
};

export const Ranking = () => {
  const currentTerm = getCurrentYearMonth();
  const colors = useColorModeValue(
    ["red.50", "blue.50"],
    ["red.900", "blue.900"]
  );
  const [tabIndex, setTabIndex] = useState(0);
  const [term, setTerm] = useState(currentTerm);
  const [displayTerm, setDisplayTerm] = useState(currentTerm);
  const [getRankList, setGetRankList] = useState<GetRanking200ResponseInner[]>(
    []
  );
  const bg = colors[tabIndex];
  const fetchRankingList = async () => {
    try {
      const response = await rankingApi.getRanking(term);
      console.log("Fetched data:", response.data);
      if (Array.isArray(response.data) && response.data.length > 0) {
        setGetRankList(response.data);
      } else {
        console.error("Invalid data structure:", response.data);
        setGetRankList([]);
      }
    } catch (error) {
      console.error("Error fetching ranking data:", error);
      setGetRankList([]);
    }
  };
  const handleKeyDown = async (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      const input = displayTerm;
      const regex = /^(\d{4}(-(?:0[1-9]|1[0-2]))?)?$/;
      if (regex.test(input)) {
        setTerm(input);
      } else {
        setDisplayTerm(term);
      }
    }
  };
  const adjustDate = async (amount: number, unit: "month" | "year") => {
    const date = new Date(displayTerm + "-01");
    let newTerm: string;
    if (unit === "month") {
      date.setMonth(date.getMonth() + amount);
      newTerm = date.toISOString().slice(0, 7);
    } else if (unit === "year") {
      date.setFullYear(date.getFullYear() + amount);
      newTerm = date.getFullYear().toString();
    } else {
      console.error("Invalid unit specified");
      return;
    }
    setDisplayTerm(newTerm);
    setTerm(newTerm);
  };
  useEffect(() => {
    fetchRankingList();
  }, [term]);
  console.log(getRankList);
  return (
    <Tabs
      isFitted
      onChange={(index) => setTabIndex(index)}
      bg={bg}
      border={"2px"}
      borderColor="gray.200"
      mt={{ base: 20, md: 0 }}
      width="100%"
    >
      <Flex
        direction={{ base: "column", md: "row" }}
        alignItems={{ base: "stretch", md: "center" }}
        width="100%"
      >
        <TabList
          mb={{ base: 4, md: "1em" }}
          width={{ base: "100%", md: "50%" }}
        >
          <Tab>滞在時間</Tab>
          <Tab>出席日数</Tab>
        </TabList>
        <Spacer display={{ base: "none", md: "block" }} />
        <Flex
          justifyContent="center"
          width={{ base: "100%", md: "50%" }}
          mb={{ base: 4, md: 2 }}
        >
          <InputGroup size="md" width="auto">
            <InputLeftAddon backgroundColor="teal.200">
              <IconButton
                backgroundColor="teal.200"
                icon={<TriangleDownIcon />}
                onClick={() => adjustDate(-1, "year")}
                aria-label="Decrease year"
                size={{ base: "sm", md: "sm" }}
              />
              <IconButton
                backgroundColor="teal.200"
                icon={<ChevronLeftIcon />}
                onClick={() => adjustDate(-1, "month")}
                aria-label="Decrease month"
                size={{ base: "sm", md: "sm" }}
                ml={1}
              />
            </InputLeftAddon>
            <Input
              value={displayTerm}
              placeholder="YYYY-MM"
              onChange={(e) => setDisplayTerm(e.target.value)}
              onKeyDown={handleKeyDown}
              maxLength={7}
              width="36"
            />
            <InputRightAddon backgroundColor="teal.200">
              <IconButton
                backgroundColor="teal.200"
                icon={<ChevronRightIcon />}
                onClick={() => adjustDate(1, "month")}
                aria-label="Increase month"
                size={{ base: "sm", md: "sm" }}
              />
              <IconButton
                backgroundColor="teal.200"
                icon={<TriangleUpIcon />}
                onClick={() => adjustDate(1, "year")}
                aria-label="Increase year"
                size={{ base: "sm", md: "sm" }}
                ml={1}
              />
            </InputRightAddon>
          </InputGroup>
        </Flex>
      </Flex>
      <TabPanels>
        {getRankList.length > 0 ? (
          <>
            <TabPanel key={"stay_time"}>
              <RankingList
                getRankList={getRankList}
                placeholder="stay_time"
              ></RankingList>
            </TabPanel>
          </>
        ) : (
          <Loading loadingItemText="ランキング情報を読込中"></Loading>
        )}
        {getRankList.length > 0 ? (
          <>
            <TabPanel key={"attendance_days"}>
              <RankingList
                getRankList={getRankList}
                placeholder="attendance_days"
              ></RankingList>
            </TabPanel>
          </>
        ) : (
          <Loading loadingItemText="ランキング情報を読込中"></Loading>
        )}
      </TabPanels>
    </Tabs>
  );
};
