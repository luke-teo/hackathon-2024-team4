import { type MouseEvent, useContext } from "react";
import { SelectedUserContext } from "./context/SelectedUserContext";
import { Box, Chip, IconButton, ListItem, Typography } from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";
import PersonIcon from "@mui/icons-material/Person";
import type { User } from "../utils/types";
import { colors } from "../utils/colors";
import { UserAvatar } from "./widgets/UserSideMenu";
import { useGetScoresByUserIdQuery } from "../services/api/v1";
import { DateTime } from "luxon";

interface Props {
	user: User;
	setValue: (v: User | null) => void;
}

export const UserListItem = ({ user, setValue }: Props) => {
	const { selectedUser, setSelectedUser } = useContext(SelectedUserContext);

	const now = DateTime.now();
	const start = now.startOf("month");
	const end = now.endOf("month");

	const { data: getUserScores } = useGetScoresByUserIdQuery({
		userId: user.id.toString() ?? "",
		startDate: start.toISODate() ?? "",
		endDate: end.toISODate() ?? "",
	});

	const scores = getUserScores?.scores;
	const latestScore =
		scores && scores.length > 1 ? scores[scores.length - 1] : null;

	const isWarning = () => {
		const zScore = Math.abs(latestScore?.zScore ?? 0);
		return zScore > 2 && zScore < 3;
	};

	const isUrgent = () => {
		const zScore = Math.abs(latestScore?.zScore ?? 0);
		return zScore > 3;
	};

	const handleSelectUser = (u: User) => {
		setSelectedUser(u);
		setValue(u);
	};

	const handleDeleteUser = (e: MouseEvent) => {
		e.preventDefault();
		setSelectedUser(null);
		setValue(null);
	};
	return (
		<ListItem
			key={user.id}
			secondaryAction={
				<>
					{selectedUser?.id === user.id ? (
						<IconButton onClick={(e) => handleDeleteUser(e)}>
							<CloseIcon />
						</IconButton>
					) : (
						<IconButton onClick={() => handleSelectUser(user)}>
							<PersonIcon
								sx={{
									color:
										selectedUser?.id === user.id
											? colors.IconSelected
											: undefined,
								}}
							/>
						</IconButton>
					)}
				</>
			}
			sx={{
				backgroundColor:
					selectedUser?.id === user.id ? colors.BackgroundSelected : undefined,
			}}
		>
			<Box
				sx={{
					display: "flex",
					alignItems: "center",
					justifyContent: "space-between",
				}}
			>
				<Box
					sx={{
						display: "flex",
						alignItems: "center",
						gap: 1,
					}}
				>
					<UserAvatar name={user.name} />

					<Typography
						sx={{
							fontSize: 14,
							fontWeight: "bold",
						}}
					>
						{user.name}
					</Typography>

					{isWarning() ? (
						<Chip color="warning" label="Review" size="small" />
					) : isUrgent() ? (
						<Chip color="error" label="Required" size="small" />
					) : (
						<></>
					)}
				</Box>
			</Box>
		</ListItem>
	);
};
