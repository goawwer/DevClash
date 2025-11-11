import AuthorCard from "./AuthorCard";
import styles from "./Authors.module.scss";

const authorsArray = [
	{ Fname: "Новоселов", Lname: "Артем", role: "Фронтенд-разработчик" },
	{ Fname: "Сельков", Lname: "Вадим", role: "Бекенд-разработчик" },
	{ Fname: "Винокуров", Lname: "Андрей", role: "Старший тестировщик" },
	{ Fname: "Фалалеев", Lname: "Евгений", role: "Проектный менеджер" },
	{ Fname: "Наговицын", Lname: "Дмитрий", role: "Младший тестировщик" },
];

export default function Authors() {
	return (
		<div className={styles.authors}>
			<h2 className={styles.authors__title}>Авторы</h2>

			<div className={styles.authors__container}>
				{authorsArray.map((item, id) => {
					return (
						<AuthorCard
							key={id}
							first_name={item.Fname}
							last_name={item.Lname}
							role={item.role}
						/>
					);
				})}
			</div>
		</div>
	);
}
