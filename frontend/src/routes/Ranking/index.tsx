import {
  Avatar,
  AvatarBadge,
  Button,
  Flex,
  Grid,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Table,
  TableContainer,
  Tabs,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
  useColorModeValue,
} from "@chakra-ui/react";
import { useState } from "react";
import { RankingList } from "../../features/RankingList";

export const Ranking = () => {
  const colors = useColorModeValue(
    ["red.50", "teal.50", "blue.50"],
    ["red.900", "teal.900", "blue.900"]
  );
  const [tabIndex, setTabIndex] = useState(0);
  const bg = colors[tabIndex];
  return (
    <Tabs
      isFitted
      onChange={(index) => setTabIndex(index)}
      bg={bg}
      border={"2px"}
      borderColor="gray.200"
    >
      <TabList mb="1em">
        <Tab>滞在時間</Tab>
        <Tab>出席日数</Tab>
        <Tab>保有コイン</Tab>
      </TabList>
      <TabPanels>
        <TabPanel>
          <RankingList placeholder="滞在時間"></RankingList>
        </TabPanel>
        <TabPanel>
          <RankingList placeholder="出席日数"></RankingList>
        </TabPanel>
        <TabPanel>
          <RankingList placeholder="保有コイン"></RankingList>
        </TabPanel>
      </TabPanels>
    </Tabs>
  );
};
