import {
  Carousel,
  CarouselApi,
  CarouselContent,
  CarouselItem,
} from "@/components/carousel";
import { ConventionsData } from "@/example";
import { Item } from "./example-item.client";
type ExampleCarouselProps = {
  api: (api: CarouselApi) => void;
  scrollRef: React.RefObject<HTMLDivElement | null>;
};
export function ExampleCarousel({ api, scrollRef }: ExampleCarouselProps) {
  return (
    <Carousel
      setApi={api}
      ref={scrollRef}
      opts={{ loop: true, watchDrag: false }}
      className="w-full scroll-mt-[80px]"
    >
      <CarouselContent>
        {[
          ConventionsData.map((data, index) => {
            return (
              <CarouselItem
                key={index}
                className="flex flex-col h-fit  w-full  gap-12  items-start justify-start  py-4"
              >
                <div className="flex h-fit w-fit flex-col gap-4">
                  <h3 className="text-2xl font-bold tracking-tight text-white">
                    {data.name}
                  </h3>
                  <p className="text-sm text-brand-offWhite">
                    {data.description}
                  </p>
                </div>
                {data.items.map((item, index) => {
                  return (
                    <Item
                      key={index}
                      index={index + 1}
                      name={item.name}
                      args={item.args}
                      description={item.description}
                      infoLink={item.infoLink}
                    />
                  );
                })}
              </CarouselItem>
            );
          }),
        ]}
      </CarouselContent>
    </Carousel>
  );
}
