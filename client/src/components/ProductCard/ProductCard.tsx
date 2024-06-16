import Image from 'next/image'

import { Product } from '@/interfaces/Product'
import { DEFAULT_PRICE_VALUE } from '@/constants/Product'

type Props = {
	product: Product;
};

export default function ProductCard({ product }: Props): JSX.Element | null {
	return (
		<div className="flex flex-col product-card border-2 border-stone-100 border-solid p-6 rounded-xl w-2/12 items-center justify-start">
			<div className="flex relative justify-center">
				<button className="py-1 px-2 flex justify-center m-auto absolute top-4 bg-yellow-300 text-sm text-stone-900 rounded-full">
					{product?.categories[0]}
				</button>
				<Image
					alt=""
					className="flex"
					src={product?.pictures ? product?.pictures[0]?.uri : ''}
					width={220}
					height={300}
					quality={100}
					priority
				/>
			</div>
			<span className="flex font-regular pt-4 h-20 text-center text-sm font-semibold">
				<a
					href="#"
					className="hover:text-blue-600"
				>
					{product?.label}
				</a>
			</span>
			<span className="text-xs">{product?.editors[0]}</span>
			<span className="price font-bold pt-1 pb-2">
				{`$${product?.price}` ?? DEFAULT_PRICE_VALUE}
			</span>
			<button className="text-white bg-black hover:bg-blue-500 rounded-full pt-2 pb-2 pl-4 pr-4 text-sm font-semibold">
				Add to bag
			</button>
		</div>
	);
}
