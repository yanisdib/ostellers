export type Product = {
	readonly id: string;
	reference: string;
	label: string;
	description?: string;
	categories: string[];
	tags?: string[];
	artists: string[];
	editors: string[];
	pictures?: Picture[];
	stock: number;
	price?: number;
	availability: ProductAvailability;
	formats: ProductFormat[];
	releasedAt: string;
	createdAt?: string;
	updatedAt?: string;
};

export type Picture = {
	uri: string;
	title: string;
	order: number;
	caption?: string;
};

export enum ProductFormat {
	PHYSICAL = 'physical',
	DIGITAL = 'digital',
}

export enum ProductAvailability {
	IN_STOCK = 'in stock',
	SOLD_OUT = 'sold out',
	RESTOCKING = 'restocking',
	LIMITED = 'limited',
	UPCOMING_RELEASE = 'upcoming release',
}

export type SimpleProduct = {
	readonly id: string;
	readonly label: string;
	readonly categories: string[];
	readonly tags?: string[];
	readonly artists: string[];
	readonly editors: string[];
	readonly pictures?: Picture[];
	readonly stock: number;
	readonly price?: number;
	readonly formats: ProductFormat[];
	readonly releasedAt: string;
};
