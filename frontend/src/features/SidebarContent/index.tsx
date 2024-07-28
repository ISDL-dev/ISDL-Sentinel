import React from "react";
import { Box, Flex, useColorModeValue, BoxProps } from "@chakra-ui/react";
import { CloseButton } from "@chakra-ui/react";
import NavItem from "../NavItem";
import { IconType } from "react-icons";
import { FiHome, FiBarChart2, FiMapPin, FiSettings } from "react-icons/fi";
import { FaHistory } from "react-icons/fa";
import { GiPerspectiveDiceSixFacesRandom } from "react-icons/gi";

interface SidebarProps extends BoxProps {
  onClose: () => void;
}

const LinkItems: Array<{ name: string; icon: IconType; href: string }> = [
  { name: "Attendee List", icon: FiHome, href: "/" },
  { name: "Access History", icon: FaHistory, href: "/access-history" },
  { name: "Ranking", icon: FiBarChart2, href: "/" },
  { name: "Gacha", icon: GiPerspectiveDiceSixFacesRandom, href: "/" },
  { name: "ISDL Map", icon: FiMapPin, href: "/" },
  { name: "Settings", icon: FiSettings, href: "/" },
];

const SidebarContent = ({ onClose, ...rest }: SidebarProps) => {
  return (
    <Box
      transition="3s ease"
      bg={useColorModeValue("white", "gray.900")}
      borderRight="1px"
      borderRightColor={useColorModeValue("gray.200", "gray.700")}
      w={{ base: "full", md: 60 }}
      pos="fixed"
      h="full"
      zIndex="0" // SidebarContentを前面に表示
      {...rest}
    >
      <Flex h="6" alignItems="center" mx="8" justifyContent="space-between">
        <CloseButton display={{ base: "flex", md: "none" }} onClick={onClose} />
      </Flex>
      {LinkItems.map((link) => (
        <NavItem key={link.name} icon={link.icon} href={link.href}>
          {link.name}
        </NavItem>
      ))}
    </Box>
  );
};

export default SidebarContent;