import { createFileRoute } from '@tanstack/react-router'
import {
    arrayMove,
    rectSortingStrategy,
    SortableContext,
    sortableKeyboardCoordinates,
    useSortable,
} from '@dnd-kit/sortable'
import {
    closestCenter,
    DndContext,
    KeyboardSensor,
    PointerSensor,
    useSensor,
    useSensors,
} from '@dnd-kit/core'
import { useEffect, useState } from 'react'
import { useSuspenseQuery } from '@tanstack/react-query'

export const Route = createFileRoute(
    '/_admin/admin/(product)/product/$productId/_details/edit'
)({
    component: EditProductComponent,
})

function SortableItem(props: { id: string; order: number; url: string }) {
    const { attributes, listeners, setNodeRef, transform, transition } =
        useSortable({ id: props.id })

    const style = {
        transform: transform
            ? `translate3d(${transform.x}px, ${transform.y}px, 0)`
            : undefined,
        transition,
    }

    return (
        <div ref={setNodeRef} style={style} {...attributes} {...listeners}>
            <div className={'h-[50px] w-[50px] bg-accent text-primary'}>
                {props.id}
            </div>
        </div>
    )
}

function SortableList({
    items,
}: {
    items: { id: string; order: number; url: string }[]
}) {
    return (
        <SortableContext items={items} strategy={rectSortingStrategy}>
            <div className={'grid grid-cols-3 gap-4'}>
                {items.map((item) => (
                    <SortableItem
                        order={item.order}
                        id={item.id}
                        url={item.url}
                    />
                ))}
            </div>
        </SortableContext>
    )
}

function EditProductComponent() {
    const { viewProductImagesOptions } = Route.useRouteContext()
    const productImagesQuery = useSuspenseQuery(viewProductImagesOptions)
    const images = productImagesQuery.data
    const [items, setItems] = useState<
        { id: string; order: number; url: string }[]
    >([])
    const sensors = useSensors(
        useSensor(PointerSensor),
        useSensor(KeyboardSensor, {
            coordinateGetter: sortableKeyboardCoordinates,
        })
    )

    useEffect(() => {
        if (images) {
            setItems(images)
        }
    }, [images])

    return (
        <DndContext sensors={sensors} collisionDetection={closestCenter}>
            <SortableList items={items} />
        </DndContext>
    )
}
